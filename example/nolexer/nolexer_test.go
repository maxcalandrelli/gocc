package nolexer

import (
	"testing"

	"github.com/maxcalandrelli/gocc/example/nolexer/internal/parser"
	"github.com/maxcalandrelli/gocc/example/nolexer/scanner"
)

func Test1(t *testing.T) {
	S := scanner.NewString("hiya world")
	P := parser.NewParser()
	if _, e, _ := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}

func Test2(t *testing.T) {
	S := scanner.NewString("hello world")
	P := parser.NewParser()
	if _, e, _ := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}
