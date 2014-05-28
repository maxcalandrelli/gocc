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

package golang

import (
	"bytes"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/lexer/items"
	"fmt"
	"path"
	"text/template"
)

func genTransitionTable(pkg, outDir, header string, itemSets *items.ItemSets) {
	fname := path.Join(outDir, "lexer", "transitiontable.go")
	io.WriteFile(fname, getTransitionTable(itemSets, header))
}

func getTransitionTable(itemsets *items.ItemSets, header string) []byte {
	tmpl, err := template.New("lexer transition table").Parse(transTabSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	data := &transitionTableData{
		Header: header,
		Rows:   getTransitionTableData(itemsets),
	}
	err = tmpl.Execute(wr, data)
	if err != nil {
		panic(err)
	}
	return wr.Bytes()
}

type transitionTableData struct {
	Header string
	Rows   []transitionTableRowData
}

func getTransitionTableData(itemsets *items.ItemSets) []transitionTableRowData {
	data := make([]transitionTableRowData, itemsets.Size())
	for setNo, set := range itemsets.List() {
		if set.SymbolClasses.MatchAny {
			data[setNo].MatchAny = true
			data[setNo].MatchAnyState = set.DotTransition
		} else {
			data[setNo].MatchAnyState = -1
		}
		data[setNo].SymRange = make([]transitionTableSymRangeData, set.SymbolClasses.Size())
		for rngNo, rng := range set.SymbolClasses.List() {
			data[setNo].SymRange[rngNo].Range = rng.String()
			data[setNo].SymRange[rngNo].Test = rangeTest(rng)
			data[setNo].SymRange[rngNo].State = set.Transitions[rngNo]
		}
		data[setNo].Imports = make([]importTransitionData, len(set.Symbols.ImportIdList))
		for impI, imprt := range set.Symbols.ImportIdList {
			if nextSet := set.ImportTransitions[impI]; nextSet != -1 {
				data[setNo].Imports[impI].Import = imprt
				data[setNo].Imports[impI].ExtFunc = set.Symbols.ExternalFunction(imprt)
				data[setNo].Imports[impI].State = nextSet
			}
		}
	}
	return data
}

func rangeTest(rng items.CharRange) string {
	if rng.From == rng.To {
		return fmt.Sprintf("r == %d", rng.From)
	}
	return fmt.Sprintf("%d <= r && r <= %d", rng.From, rng.To)
}

type transitionTableRowData struct {
	MatchAny      bool
	MatchAnyState int
	SymRange      []transitionTableSymRangeData
	Imports       []importTransitionData
}

type importTransitionData struct {
	Import  string
	ExtFunc string
	State   int
}

type transitionTableSymRangeData struct {
	Range string
	Test  string
	State int
}

const transTabSrc = `
package lexer

{{.Header}}

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	{{range $sno, $state := .Rows}}
		// S{{$sno}}
		func(r rune) int {
			switch {
			{{range $rng := $state.SymRange}}case {{$rng.Test}} : // {{$rng.Range}}
				return {{$rng.State}}
			{{end}}
			{{range $imp := $state.Imports}}{{if $imp.State}}case {{$imp.ExtFunc}}(r): // {{$imp.Import}}
				return {{$imp.State}}
			{{end}}{{end}}
			{{if $state.MatchAny}}default:
				return {{$state.MatchAnyState}}
			}
			{{else}}
			}
			return NoState
			{{end}}
		},
	{{end}}
}
`
