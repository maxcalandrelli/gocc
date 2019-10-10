package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	rr "github.com/maxcalandrelli/gocc/example/rr"
)

func showResult(r interface{}, e error, l int) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
	} else {
		fmt.Printf("r=%#v, %d bytes\n", r, l)
	}
}

var (
	File    string
	Text    string
	Longest bool
)

func parse(longest bool, lex *rr.Lexer) (res interface{}, err error, ptl int) {
	if longest {
		return rr.NewParser().ParseLongestPrefix(lex)
	} else {
		return rr.NewParser().Parse(lex)
	}
	return
}

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(parse(Longest, rr.NewLexerString(Text)))
	}
	if File > "" {
		l, e := rr.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, rr.NewLexerString(str)))
	}
}
