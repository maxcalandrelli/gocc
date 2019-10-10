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
	"path"
	"text/template"

	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/io"
	"github.com/maxcalandrelli/gocc/internal/lexer/items"
)

func genLexer(pkg, outDir string, itemsets *items.ItemSets, cfg config.Config, subpath string) {
	tmpl, err := template.New("lexer").Parse(lexerSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, getLexerData(pkg, outDir, itemsets, cfg, subpath))
	if err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(outDir, subpath, "lexer", "lexer.go"), buf.Bytes())
}

func getLexerData(pkg, outDir string, itemsets *items.ItemSets, cfg config.Config, subpath string) *lexerData {
	lexSymbols := itemsets.Symbols().List()
	return &lexerData{
		Debug:      cfg.DebugLexer(),
		PkgPath:    pkg,
		SubPath:    subpath,
		NumStates:  itemsets.Size(),
		NumSymbols: len(lexSymbols),
		Symbols:    lexSymbols,
	}
}

type lexerData struct {
	Debug      bool
	PkgPath    string
	SubPath    string
	NumStates  int
	NumSymbols int
	NextState  []byte
	Symbols    []string
}

const lexerSrc string = `
// Code generated by gocc; DO NOT EDIT.

package lexer

import (
  {{if .Debug}}	"fmt" {{end}}
  "io"
  "bytes"
  "os"

  {{if .Debug}}	"{{.PkgPath}}/{{.SubPath}}/util" {{end}}
  "{{.PkgPath}}/iface"
  "{{.PkgPath}}/{{.SubPath}}/token"
  "{{.PkgPath}}/{{.SubPath}}/io/stream"
)

const (
	NoState    = -1
	NumStates  = {{.NumStates}}
	NumSymbols = {{.NumSymbols}}
  INVALID_RUNE = rune(-1)
)

type position struct {
	token.Pos
}

type Lexer struct {
	position
	stream   iface.TokenStream
	eof      bool
}

func NewLexerBytes(src []byte) *Lexer {
	lexer := &Lexer{stream: bytes.NewReader(src)}
	lexer.reset()
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	s, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	lexer := &Lexer{stream: stream.NewWindowReader(s)}
	lexer.reset()
	return lexer, nil
}

func NewLexer(reader io.Reader) (*Lexer, error) {
	lexer := &Lexer{}
	lexer.reset()
	if lexer.stream, _ = reader.(iface.TokenStream); lexer.stream == nil {
		lexer.stream = stream.NewWindowReader(reader)
	}
	return lexer, nil
}

func (l *Lexer) reset () {
  l.position.Reset()
}

func (l Lexer) GetStream() iface.TokenStream {
  return l.stream
}

type checkPoint int64

func (c checkPoint) DistanceFrom (o iface.CheckPoint) int {
  return int (c - o.(checkPoint))
}

func (l Lexer) GetCheckPoint() iface.CheckPoint {
  pos, _ := l.stream.Seek(0, io.SeekCurrent)
  return checkPoint(pos)
}

func (l Lexer) GotoCheckPoint(cp iface.CheckPoint) {
  l.stream.Seek(int64(cp.(checkPoint)), io.SeekStart)
}

func (l *Lexer) Scan() (tok *token.Token) {
	{{- if .Debug}}
	fmt.Printf("Lexer.Scan() pos=%d\n", l.position.Pos.Offset)
	{{- end}}
	tok = new(token.Token)
	tok.Type = token.INVALID
	tok.Lit = []byte{}
  start := l.position
  state := 0
	for state != -1  {
		{{- if .Debug}}
		fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", l.position.Pos.Offset, l.position.Pos.Line, l.position.Pos.Column, state)
		{{- end}}
	  curr, size, err := l.stream.ReadRune()
    if size < 1 || err != nil {
      curr = INVALID_RUNE
    }
		{{- if .Debug}}
		fmt.Printf("\trune=<%c> size=%d err=%v\n", curr, size, err)
		{{- end}}
		if size > 0 {
  		l.position.Pos.Offset += size
    }
		nextState := -1
    if err == nil {
  		if curr != INVALID_RUNE {
  			nextState = TransTab[state](curr)
  		}
  		{{- if .Debug}}
  		fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(curr), curr, nextState)
  		fmt.Printf("\t\tpos=%d, size=%d, start=%d\n", l.position.Pos.Offset, size, start.Pos.Offset)
  		if nextState != -1 {
  			fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
  		}
  		{{- end}}
    }
		state = nextState
		if state != -1 {
    	switch curr {
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
				tok.Lit = append(tok.Lit, string(curr)...)
			case ActTab[state].Ignore != "":
				start = l.position
				state = 0
				tok.Lit = []byte{}
			}
		} else {
      l.stream.UnreadRune()
    }
  	if err == io.EOF && len(tok.Lit)==0 {
  		tok.Type = token.EOF
  		tok.Pos = start.Pos
  		return
  	}
	}
	tok.Pos = start.Pos
	{{- if .Debug}}
	fmt.Printf("Token at %s: %s \"%s\"\n", tok.String(), token.TokMap.Id(tok.Type), tok.Lit)
	{{- end}}
	return
}

func (l *Lexer) Reset() {
	l.position.Reset()
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
`
