// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/maxcalandrelli/gocc/example/rr/iface"
	parseError "github.com/maxcalandrelli/gocc/example/rr/internal/errors"
	"github.com/maxcalandrelli/gocc/example/rr/internal/token"
)

const (
	numProductions = 7
	numStates      = 7
	numSymbols     = 8
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

type TokenStream = iface.TokenStream

func NewParser() *Parser {
	return NewParserWithContext(nil)
}

func NewParserWithContext(u interface{}) *Parser {
	p := &Parser{stack: newStack(), userContext: u}
	p.Reset()
	return p
}

func (p *Parser) Reset() {
	p.stack.reset()
	p.stack.push(0, nil)
}

func (p *Parser) Error(err error, scanner iface.Scanner) (recovered bool, errorAttrib *parseError.Error) {
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

func (p *Parser) Parse(scanner iface.Scanner) (res interface{}, err error, ptl int) {
	return p.parse(scanner, false)
}

func (p *Parser) ParseLongestPrefix(scanner iface.Scanner) (res interface{}, err error, ptl int) {
	return p.parse(scanner, true)
}

var errNotRepositionable = errors.New("scanner not repositionable")

func (p *Parser) parse(scanner iface.Scanner, longest bool) (res interface{}, err error, ptl int) {
	var (
		tokens     iface.CheckPointable
		afterPos   iface.CheckPoint
		checkPoint iface.CheckPoint
	)
	readNextToken := func() error {
		if tokens == nil && (len(parserActions.table[p.stack.top()].cdActions) > 0 || longest) {
			return errNotRepositionable
		}
		checkPoint = tokens.GetCheckPoint()
		p.nextToken = scanner.Scan()
		if longest {
			afterPos = tokens.GetCheckPoint()
		}
		return nil
	}
	p.Reset()
	tokens, _ = scanner.(iface.CheckPointable)
	startCp := tokens.GetCheckPoint()
	if err := readNextToken(); err != nil {
		return nil, err, tokens.GetCheckPoint().DistanceFrom(startCp)
	}
	for acc := false; !acc; {
		action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]
		if action == nil {
			//
			// If no action, check if we have some context dependent parsing to try
			//
			for _, cdAction := range parserActions.table[p.stack.top()].cdActions {
				tokens.GotoCheckPoint(checkPoint)
				cd_res, cd_err, cd_parsed := cdAction.tokenScanner(scanner.GetStream(), p.userContext)
				if cd_err == nil && len(cd_parsed) > 0 {
					action = parserActions.table[p.stack.top()].actions[cdAction.tokenIndex]
					if action != nil {
						p.nextToken.Foreign = true
						p.nextToken.ForeignAstNode = cd_res
						p.nextToken.Lit = cd_parsed
						p.nextToken.Type = token.Type(cdAction.tokenIndex)
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
			tokens.GotoCheckPoint(checkPoint)
			action = parserActions.table[p.stack.top()].actions[token.EOF]
			if action == nil {
				//
				//  ok, let's consume the token then
				//
				tokens.GotoCheckPoint(afterPos)
			}
		}

		//
		//  Well, no action is no action after all...
		//
		if action == nil {
			if recovered, errAttrib := p.Error(nil, scanner); !recovered {
				p.nextToken = errAttrib.ErrorToken
				return nil, p.newError(nil), tokens.GetCheckPoint().DistanceFrom(startCp)
			}
			if action = parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action == nil {
				panic("Error recovery led to invalid action")
			}
		}

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
				return nil, err, tokens.GetCheckPoint().DistanceFrom(startCp)
			}
		case reduce:
			prod := productionsTable[int(act)]
			attrib, err := prod.ReduceFunc(p.userContext, p.stack.popN(prod.NumSymbols))
			if err != nil {
				return nil, p.newError(err), tokens.GetCheckPoint().DistanceFrom(startCp)
			} else {
				p.stack.push(gotoTab[p.stack.top()][prod.NTType], attrib)
			}
		default:
			panic("unknown action: " + action.String())
		}
	}
	return res, nil, tokens.GetCheckPoint().DistanceFrom(startCp)
}
