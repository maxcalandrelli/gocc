// Code generated by gocc; DO NOT EDIT.

package lexer

import (

  "io"
  "bytes"
  "os"


"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/token"
"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/io/stream"
)

const (
	NoState    = -1
	NumStates  = 114
	NumSymbols = 79
)

type position struct {
	token.Pos
	StreamPosition int64
}

type Lexer struct {
	position
	stream token.TokenStream
	eof    bool
}

func NewLexerBytes(src []byte) *Lexer {
	lexer := &Lexer{stream: bytes.NewReader(src)}
	lexer.position.Reset()
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	s, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	lexer := &Lexer{stream: stream.NewWindowReader(s)}
	lexer.position.Reset()
	return lexer, nil
}

func NewLexer(reader io.Reader) (*Lexer, error) {
	lexer := &Lexer{}
	lexer.position.Reset()
	if lexer.stream, _ = reader.(token.TokenStream); lexer.stream == nil {
		lexer.stream = stream.NewWindowReader(reader)
	} else {
		lexer.position.StreamPosition, _ = lexer.stream.Seek(0, io.SeekCurrent)
	}
	return lexer, nil
}

func (l Lexer) GetStream() io.Reader {
  return l.stream
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.eof {
		tok.Type = token.EOF
		tok.Pos = l.position.Pos
		return
	}
	l.position.StreamPosition, _ = l.stream.Seek(0, io.SeekCurrent)
	start, end := l.position, position{}
	tok.Type = token.INVALID
	tok.Lit = []byte{}
	state, rune1 := 0, rune(-1)
	for state != -1 {
		if l.eof {
			rune1 = -1
		} else {
			rune2, size, err := l.stream.ReadRune()
			if err == io.EOF {
				l.eof = true
				err = nil
			}
			if err == nil && size > 0 {
				rune1 = rune2
				l.position.StreamPosition += int64(size)
				l.position.Pos.Offset += size
			}
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {
			switch rune1 {
			case '\n':
				l.position.Pos.Line++
				l.position.Pos.Column = 1
			case '\r':
				l.position.Pos.Column = 1
			case '\t':
				l.position.Pos.Column += 4
			default:
				l.position.Pos.Column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				l.position.StreamPosition, _ = l.stream.Seek(0, io.SeekCurrent)
				end = l.position
				tok.Lit = append(tok.Lit, string(rune1)...)
			case ActTab[state].Ignore != "":
				start = l.position
				state = 0
				tok.Lit = []byte{}
				if l.eof {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.position
			}
		}
	}
	if end.Pos.Offset > start.Pos.Offset {
		l.Reposition(end)
	}
	tok.Pos = start.Pos
	return
}

func (l *Lexer) Reset() {
	l.position.Reset()
}

func (l *Lexer) Reposition(p position) {
	l.position = p
	l.stream.Seek(l.position.StreamPosition, io.SeekStart)
}

func (l Lexer) CurrentPosition() position {
	return l.position
}

func (p *position) Reset() {
	p.Offset = 0
	p.Line = 1
	p.Column = 1
}

func (p position) StartingFrom(base position) position {
	r := p
	r.Pos = p.Pos.StartingFrom(base.Pos)
	return r
}
