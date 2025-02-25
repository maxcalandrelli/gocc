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
	"github.com/maxcalandrelli/gocc/internal/config"
)

type SyntaxError struct {
	string
	StdSyntaxSymbol
}

var ErrorSymbol = SyntaxError{config.SYMBOL_ERROR, StdSyntaxSymbol{}}

func (SyntaxError) SymbolString() string {
	return ErrorSymbol.string
}

func (SyntaxError) String() string {
	return ErrorSymbol.string
}

func (SyntaxError) IsError() bool { return true }

func (SyntaxError) IsTerminal() bool { return true }

func (SyntaxError) SymbolName() string {
	return config.INTERNAL_SYMBOL_ERROR
}
