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

package ast

import (
	"fmt"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/config"
)

type SyntaxSymbols []SyntaxSymbol

func NewSyntaxSymbols(sym interface{}) (SyntaxSymbols, error) {
	return SyntaxSymbols{sym.(SyntaxSymbol)}, nil
}

func AddSyntaxSymbol(symbols, symbol interface{}) (SyntaxSymbols, error) {
	return append(symbols.(SyntaxSymbols), symbol.(SyntaxSymbol)), nil
}

func (this SyntaxSymbols) String() string {
	w := new(strings.Builder)
	for i, sym := range this {
		if i > 0 {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, sym.SymbolName())
	}
	return w.String()
}

func NewSyntaxSymbolsFromToken(tok interface{}) (SyntaxSymbols, error) {
	switch getString(tok) {
	case config.SYMBOL_EMPTY:
		return SyntaxSymbols{EmptySymbol}, nil
	case config.SYMBOL_ERROR:
		return SyntaxSymbols{ErrorSymbol}, nil
	}
	sym, err := NewStringLit(tok)
	fmt.Printf("LITERAL:<%q>\n", sym)
	return SyntaxSymbols{sym}, err
}

type SyntaxSymbolsByName SyntaxSymbols

func (s *SyntaxSymbolsByName) Len() int {
	return len(*s)
}

func (s *SyntaxSymbolsByName) Less(i, j int) bool {
	return (*s)[i].SymbolString() < (*s)[j].SymbolString()
}

func (s *SyntaxSymbolsByName) Swap(i, j int) {
	t := (*s)[i]
	(*s)[i] = (*s)[j]
	(*s)[j] = t
}
