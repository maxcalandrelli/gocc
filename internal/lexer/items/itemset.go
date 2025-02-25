//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package items

import (
	"fmt"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/ast"
	"github.com/maxcalandrelli/gocc/internal/lexer/symbols"
)

/*
Set of basic lexical items
Each entry of Transitions is the next set for the corresponding symbo class.
len(ItemSet.Transitions) == ItemSet.SymbolClasses.Size()
*/
type ItemSet struct {
	setNo   int
	Items   ItemList
	lexPart *ast.LexPart
	*symbols.Symbols
	SymbolClasses     *DisjunctRangeSet
	Transitions       []int
	ImportTransitions []int
	DotTransition     int
}

func NewItemSet(setNo int, lexPart *ast.LexPart, symbols *symbols.Symbols, items ItemList) *ItemSet {
	if debug_deeply {
		trace(setNo, "creating, start items:\n%s", items)
	}
	set := &ItemSet{
		setNo:         setNo,
		Items:         items.Closure(lexPart, symbols),
		lexPart:       lexPart,
		Symbols:       symbols,
		DotTransition: -1,
	}
	set.getSymbolClasses(lexPart, symbols)
	set.newTransitions()
	if debug_deeply {
		trace(setNo, "created, items:\n%s\nsymbol classes:\n%s", set.Items, set.SymbolClasses)
	}
	return set
}

func ItemsSet0(lexPart *ast.LexPart, symbols *symbols.Symbols) ItemList {
	items := NewItemList(16)
	for _, p := range lexPart.ProdList.Productions {
		switch p.(type) {
		case *ast.LexTokDef, *ast.LexIgnoredTokDef:
			//TODO: shouldn't .Emoves() be  inside NewItem?
			items = items.AddNoDuplicate(NewItem(p.Id(), lexPart, symbols).Emoves()...)
		}
	}
	return items
}

func (this *ItemSet) Action() Action {
	actionItem := (*Item)(nil)
	for _, item := range this.Items {
		if !item.Prod.RegDef() && item.Reduce() {
			if actionItem == nil ||
				this.lexPart.StringLitTokDef(item.Id) != nil ||
				(this.lexPart.StringLitTokDef(actionItem.Id) == nil && item.ProdIndex < actionItem.ProdIndex) {

				actionItem = item
			}
		}
	}
	if actionItem == nil {
		return nil
	}

	switch actionItem.Prod.(type) {
	case *ast.LexTokDef:
		return Accept(actionItem.Id)
	case *ast.LexIgnoredTokDef:
		return Ignore(actionItem.Id)
	}

	return nil
}

/*
Returns the number of new items added to this.
*/
func (this *ItemSet) Add(items ...*Item) {
	this.Items.AddNoDuplicate(items...)
}

func (this *ItemSet) Contain(that *Item) bool {
	return this.Items.Contain(that)
}

func (this *ItemSet) Equal(items ItemList) bool {
	return this.Items.Equal(items)
}

func (this *ItemSet) Empty() bool {
	return len(this.Items) == 0
}

func (this *ItemSet) getSymbolClasses(lexPart *ast.LexPart, symbols *symbols.Symbols) {
	this.SymbolClasses = NewDisjunctRangeSet()
	for _, item := range this.Items {
		if !item.Reduce() {
			this.addToSymbolClasses(lexPart, symbols, item.ExpectedSymbol(), false)
		}
	}
	for _, item := range this.Items {
		if !item.Reduce() {
			this.addToSymbolClasses(lexPart, symbols, item.ExpectedSymbol(), true)
		}
	}

}

func (this *ItemSet) addToSymbolClasses(lexPart *ast.LexPart, symbols *symbols.Symbols, node ast.LexNode, negated bool) {
	addNow := !negated
	tNode := ast.LexTNode(nil)
	switch n := node.(type) {
	case *ast.LexCharLit:
		addNow = (negated == n.Negate)
		tNode = node.(ast.LexTNode)
	case *ast.LexCharRange:
		addNow = (negated == n.Negate)
		tNode = node.(ast.LexTNode)
	}
	if addNow && tNode != nil {
		this.SymbolClasses.AddLexTNode(tNode)
	}
}

func (this *ItemSet) newTransitions() {
	this.Transitions = make([]int, this.SymbolClasses.Size())
	for i := range this.Transitions {
		this.Transitions[i] = -1
	}
	this.ImportTransitions = make([]int, len(this.Symbols.ImportIdList))
	for i := range this.ImportTransitions {
		this.ImportTransitions[i] = -1
	}
}

/*
See algorithm: set.Next() in package doc
*/
func (this *ItemSet) Next(rng CharRange) ItemList {
	trace(this.setNo, "rng=%q\n", rng)
	nextItems := NewItemList(16)
	for _, item := range this.Items {
		nextItems = nextItems.AddNoDuplicate(item.Move(rng)...)
	}
	nextItems = this.dependentsClosure(nextItems)
	return nextItems.Closure(this.lexPart, this.Symbols)
}

func (this *ItemSet) NextDot() ItemList {
	nextItems := NewItemList(16)
	for _, item := range this.Items {
		nextItems = nextItems.AddNoDuplicate(item.MoveDot()...)
	}
	nextItems = this.dependentsClosure(nextItems)
	return nextItems.Closure(this.lexPart, this.Symbols)
}

func (this *ItemSet) NextImport(imprt string) ItemList {
	nextItems := NewItemList(16)
	for _, item := range this.Items {
		nextItems = nextItems.AddNoDuplicate(item.MoveRegDefId(imprt)...)
	}
	nextItems = this.dependentsClosure(nextItems)
	return nextItems.Closure(this.lexPart, this.Symbols)
}

func (this *ItemSet) dependentsClosure(items ItemList) ItemList {
	if len(items) == 0 {
		return items
	}
	//trace(this.setNo, "dependentsClosure\n    actual: %s\n    items: %s\n", util.EscapedString(this.String()).Unescape(), util.EscapedString(fmt.Sprintf("%q", items)).Unescape())
	trace(this.setNo, "dependentsClosure\n")
	for i := 0; i < len(items); i++ {
		for _, thisItem := range this.Items {
			if expSym := thisItem.ExpectedSymbol(); expSym != nil && expSym.String() == items[i].Id {
				if items[i].Reduce() {
					mv := thisItem.MoveRegDefId(items[i].Id)
					for _, mvi := range mv {
						dbg(this.setNo, "  ====> %s\n", mvi)
					}
					items = items.AddNoDuplicate(thisItem.MoveRegDefId(items[i].Id)...)
				} else {
					dbg(this.setNo, "  ==>   %s\n", thisItem)
					items = items.AddNoDuplicate(thisItem)
				}
			}
		}
	}
	return items
}

func (this *ItemSet) Size() int {
	return len(this.Items)
}

func (this *ItemSet) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "{\n")
	for _, i := range this.Items {
		fmt.Fprintf(buf, "\t%s\n", i)
	}
	fmt.Fprintf(buf, "}\n")
	fmt.Fprintf(buf, "Transitions:\n")
	for rng, setNo := range this.Transitions {
		fmt.Fprintf(buf, "\t%s -> S%d\n", this.SymbolClasses.Range(rng), setNo)
	}
	fmt.Fprintf(buf, "Symbols classes: %s\n", this.SymbolClasses)
	return buf.String()
}

func (this *ItemSet) StringItems() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "{\n")
	for _, i := range this.Items {
		fmt.Fprintf(buf, "\t%s\n", i)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}
