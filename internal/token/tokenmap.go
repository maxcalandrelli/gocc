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

package token

import (
	"github.com/maxcalandrelli/gocc/internal/ast"
)

type TokenMap struct {
	IdMap   map[string]int
	TypeMap []ast.SyntaxSymbol
	LitMap  map[int]string
}

func NewTokenMap(symbols ast.SyntaxSymbols) *TokenMap {
	tm := &TokenMap{
		IdMap:   make(map[string]int),
		TypeMap: make([]ast.SyntaxSymbol, len(symbols)),
		LitMap:  make(map[int]string),
	}

	for i, sym := range symbols {
		tm.IdMap[sym.SymbolName()] = i
		tm.TypeMap[i] = sym
		switch lit := sym.(type) {
		case ast.SyntaxStringLit:
			tm.LitMap[i] = lit.SymbolString()
		}
	}
	return tm
}
