//Copyright 2012 Vastech SA (PTY) LTD
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
	"fmt"
	"go/format"
	"path"
	"text/template"

	"github.com/maxcalandrelli/gocc/internal/ast"
	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/io"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/parser/symbols"
)

func GenParser(pkg, outDir string, prods ast.SyntaxProdList, itemSets *items.ItemSets, symbols *symbols.Symbols, cfg config.Config, subpath string) {
	tmpl, err := template.New("parser").Parse(parserSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, getParserData(pkg, subpath, prods, itemSets, symbols, cfg)); err != nil {
		panic(err)
	}
	source, err := format.Source(wr.Bytes())
	if err != nil {
		panic(fmt.Sprintf("%s in\n%s", err.Error(), wr.String()))
	}
	io.WriteFile(path.Join(outDir, subpath, "parser", "parser.go"), source)
}

type parserData struct {
	Debug          bool
	ErrorImport    string
	TokenImport    string
	NumProductions int
	NumStates      int
	NumSymbols     int
	CdTokList      ast.SyntaxSymbols
}

func getParserData(pkg, subpath string, prods ast.SyntaxProdList, itemSets *items.ItemSets, symbols *symbols.Symbols, cfg config.Config) *parserData {
	return &parserData{
		Debug:          cfg.DebugParser(),
		ErrorImport:    path.Join(pkg, subpath, "errors"),
		TokenImport:    path.Join(pkg, subpath, "token"),
		NumProductions: len(prods),
		NumStates:      itemSets.Size(),
		NumSymbols:     symbols.NumSymbols(),
		CdTokList:      symbols.ListContextDependentTokenSymbols(),
	}
}

const parserSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
  "io"
	"strings"
  "errors"

	parseError "{{.ErrorImport}}"
	"{{.TokenImport}}"
)

const (
	numProductions = {{.NumProductions}}
	numStates      = {{.NumStates}}
	numSymbols     = {{.NumSymbols}}
)

// Stack

type stack struct {
	state  []int
	attrib []Attrib
}

const INITIAL_STACK_SIZE = 100

func newStack() *stack {
	return &stack{
		state:  make([]int, 0, INITIAL_STACK_SIZE),
		attrib: make([]Attrib, 0, INITIAL_STACK_SIZE),
	}
}

func (s *stack) reset() {
	s.state = s.state[:0]
	s.attrib = s.attrib[:0]
}

func (s *stack) push(state int, a Attrib) {
	s.state = append(s.state, state)
	s.attrib = append(s.attrib, a)
}

func (s *stack) top() int {
	return s.state[len(s.state)-1]
}

func (s *stack) peek(pos int) int {
	return s.state[pos]
}

func (s *stack) topIndex() int {
	return len(s.state) - 1
}

func (s *stack) popN(items int) []Attrib {
	lo, hi := len(s.state)-items, len(s.state)

	attrib := s.attrib[lo:hi]

	s.state = s.state[:lo]
	s.attrib = s.attrib[:lo]

	return attrib
}

func (s *stack) String() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "stack:\n")
	for i, st := range s.state {
		fmt.Fprintf(w, "\t%d: %d , ", i, st)
		if s.attrib[i] == nil {
			fmt.Fprintf(w, "nil")
		} else {
			switch attr := s.attrib[i].(type) {
			case *token.Token:
				fmt.Fprintf(w, "%s", attr.Lit)
			default:
				fmt.Fprintf(w, "%v", attr)
			}
		}
		fmt.Fprintf(w, "\n")
	}
	return w.String()
}

// Parser

type Parser struct {
	stack       *stack
	nextToken   *token.Token
  userContext interface{}
}

type Scanner interface {
	Scan() (tok *token.Token)
  GetStream () io.Reader
}

type TokenStream = token.TokenStream

{{- range $c := .CdTokList }}
{{ printf "func cdFunc_%s (Stream TokenStream, Context interface{}) (interface{}, error, []byte) {return %s}" $c.SymbolString $c.ContexDependentParseFunctionCall }}
{{- end }}


func NewParser() *Parser {
	return NewParserWithContext(nil)
}

func NewParserWithContext(u interface{}) *Parser {
	p := &Parser{stack: newStack(), userContext: u }
	p.Reset()
	return p
}

func (p *Parser) Reset() {
	p.stack.reset()
	p.stack.push(0, nil)
}

func (p *Parser) Error(err error, scanner Scanner) (recovered bool, errorAttrib *parseError.Error) {
	errorAttrib = &parseError.Error{
		Err:            err,
		ErrorToken:     p.nextToken,
		ErrorSymbols:   p.popNonRecoveryStates(),
		ExpectedTokens: make([]string, 0, 8),
	}
	for t, action := range parserActions.table[p.stack.top()].actions {
		if action != nil {
			errorAttrib.ExpectedTokens = append(errorAttrib.ExpectedTokens, token.TokMap.Id(token.Type(t)))
		}
	}

	if action := parserActions.table[p.stack.top()].actions[token.TokMap.Type("error")]; action != nil {
		p.stack.push(int(action.(shift)), errorAttrib) // action can only be shift
	} else {
		return
	}

	if action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action != nil {
		recovered = true
	}
	for !recovered && p.nextToken.Type != token.EOF {
		p.nextToken = scanner.Scan()
		if action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action != nil {
			recovered = true
		}
	}

	return
}

func (p *Parser) popNonRecoveryStates() (removedAttribs []parseError.ErrorSymbol) {
	if rs, ok := p.firstRecoveryState(); ok {
		errorSymbols := p.stack.popN(p.stack.topIndex() - rs)
		removedAttribs = make([]parseError.ErrorSymbol, len(errorSymbols))
		for i, e := range errorSymbols {
			removedAttribs[i] = e
		}
	} else {
		removedAttribs = []parseError.ErrorSymbol{}
	}
	return
}

// recoveryState points to the highest state on the stack, which can recover
func (p *Parser) firstRecoveryState() (recoveryState int, canRecover bool) {
	recoveryState, canRecover = p.stack.topIndex(), parserActions.table[p.stack.top()].canRecover
	for recoveryState > 0 && !canRecover {
		recoveryState--
		canRecover = parserActions.table[p.stack.peek(recoveryState)].canRecover
	}
	return
}

func (p *Parser) newError(err error) error {
	e := &parseError.Error{
		Err:        err,
		StackTop:   p.stack.top(),
		ErrorToken: p.nextToken,
	}
	actRow := parserActions.table[p.stack.top()]
	for i, t := range actRow.actions {
		if t != nil {
			e.ExpectedTokens = append(e.ExpectedTokens, token.TokMap.Id(token.Type(i)))
		}
	}
	return e
}

func (p *Parser) Parse(scanner Scanner) (res interface{}, err error) {
  r, e, _ := p.parse(scanner, false)
  return r, e
}

func (p *Parser) ParseLongestPrefix(scanner Scanner) (res interface{}, err error, parsed []byte) {
  return p.parse(scanner, true)
}

var errNotRepositionable = errors.New("scanner not repositionable")


func (p *Parser) parse(scanner Scanner, longest bool) (res interface{}, err error, parsed []byte) {
  var (
    tokens     TokenStream
    afterPos   int64
    checkPoint int64
  )
  readNextToken := func () error {
    if tokens == nil && (len(parserActions.table[p.stack.top()].cdActions) > 0 || longest) {
      if tokens, _ = scanner.GetStream().(TokenStream); tokens == nil {
        return errNotRepositionable
      }
    }
    if len(parserActions.table[p.stack.top()].cdActions) > 0 {
      checkPoint, _ = tokens.Seek (0, io.SeekCurrent)
    }
  	p.nextToken = scanner.Scan()
    if longest {
      afterPos , _ = tokens.Seek (0, io.SeekCurrent)
    }
    return nil
  }
	p.Reset()
  if err := readNextToken(); err != nil {
    return nil, err, []byte{}
  }
	for acc := false; !acc; {
		action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]
		if action == nil {
      //
      // If no action, check if we have some context dependent parsing to try
      //
			for _, cdAction := range parserActions.table[p.stack.top()].cdActions {
				tokens.Seek(checkPoint, io.SeekStart)
				cd_res, cd_err, cd_parsed := cdAction.tokenScanner(tokens, p.userContext)
				if cd_err == nil && len(cd_parsed) > 0 {
					action = parserActions.table[p.stack.top()].actions[cdAction.tokenIndex]
          if action != nil {
            p.nextToken.Foreign = true
            p.nextToken.ForeignAstNode = cd_res
  					p.nextToken.Lit = cd_parsed
					  break
          }
				}
			}
    }
    //
    //  Still no action? If a longest possible parsing is requested in place
    //  of a full text, we should try to check if an EOF here would have led
    //  to a successful parsing instead
    //  Rember that token.EOF is 1, that is the index of SyntaxEof into symbol table
    //  Dirty, dirty, dirty... but I found it as it is, and I prefer not to touch
    //
		if action == nil && longest {
			tokens.Seek(checkPoint, io.SeekStart)
      action = parserActions.table[p.stack.top()].actions[token.EOF]
      if action == nil {
        //
        //  ok, let's consume the token then
        //
				tokens.Seek(afterPos, io.SeekStart)
      }
    }

    //
    //  Well, no action is no action after all...
    //
    if action == nil {
			if recovered, errAttrib := p.Error(nil, scanner); !recovered {
				p.nextToken = errAttrib.ErrorToken
				return nil, p.newError(nil), []byte{}
			}
			if action = parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action == nil {
				panic("Error recovery led to invalid action")
			}
    }
		{{- if .Debug }}
		fmt.Printf("S%d %s %s\n", p.stack.top(), token.TokMap.TokenString(p.nextToken), action)
		{{- end }}

		switch act := action.(type) {
		case accept:
			res = p.stack.popN(1)[0]
			acc = true
		case shift:
		  p.stack.push(int(act), p.nextToken)
      if p.nextToken.Foreign {
			  p.stack.push(int(act), p.nextToken.ForeignAstNode)
      }
      if err := readNextToken(); err != nil {
        return nil, err, []byte{}
      }
		case reduce:
			prod := productionsTable[int(act)]
			attrib, err := prod.ReduceFunc(p.userContext, p.stack.popN(prod.NumSymbols))
			if err != nil {
				return nil, p.newError(err), []byte{}
			} else {
				p.stack.push(gotoTab[p.stack.top()][prod.NTType], attrib)
			}
		default:
			panic("unknown action: " + action.String())
		}
	}
	return res, nil, []byte{}
}
`
