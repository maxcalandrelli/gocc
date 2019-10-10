// Code generated by gocc; DO NOT EDIT.

package sr

import (
	"io"

	"github.com/maxcalandrelli/gocc/example/sr/iface"
	"github.com/maxcalandrelli/gocc/example/sr/internal/io/stream"
	"github.com/maxcalandrelli/gocc/example/sr/internal/lexer"
	"github.com/maxcalandrelli/gocc/example/sr/internal/parser"
	"github.com/maxcalandrelli/gocc/example/sr/internal/token"
)

type (
	Token        = token.Token
	Lexer        = lexer.Lexer
	Parser       = parser.Parser
	TokenStream  = iface.TokenStream
	WindowReader = stream.WindowReader
	Scanner      = iface.Scanner
)

func ParseFile(fpath string) (interface{}, error, int) {
	if lexer, err := NewLexerFile(fpath); err == nil {
		return NewParser().Parse(lexer)
	} else {
		return nil, err, 0
	}
}

func ParseText(text string) (interface{}, error, int) {
	return NewParser().Parse(NewLexerBytes([]byte(text)))
}

func NewLexerBytes(src []byte) *lexer.Lexer {
	return lexer.NewLexerBytes(src)
}

func NewLexerString(src string) *lexer.Lexer {
	return lexer.NewLexerBytes([]byte(src))
}

func NewLexerFile(fpath string) (*lexer.Lexer, error) {
	return lexer.NewLexerFile(fpath)
}

func NewParser() *parser.Parser {
	return parser.NewParser()
}

func NewWindowReaderFromBytes(src []byte) WindowReader {
	return stream.NewWindowReaderFromBytes(src)
}

func NewWindowReader(rdr io.Reader) WindowReader {
	return stream.NewWindowReader(rdr)
}

func NewLimitedWindowReader(rdr io.Reader, sizeMin, sizeMax int) WindowReader {
	return stream.NewLimitedWindowReader(rdr, sizeMin, sizeMax)
}
