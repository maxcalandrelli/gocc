// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/maxcalandrelli/gocc/example/bools/ast"
)

import (
	"fmt"
	"github.com/maxcalandrelli/gocc/example/bools/bools.grammar/bools/internal/token"
	"github.com/maxcalandrelli/gocc/example/bools/bools.grammar/bools/internal/util"
	"strings"
)

func getString(X Attrib) string {
	switch X.(type) {
	case *token.Token:
		return string(X.(*token.Token).Lit)
	case string:
		return X.(string)
	}
	return fmt.Sprintf("%q", X)
}

func unescape(s string) string {
	return util.EscapedString(s).Unescape()
}

func unquote(s string) string {
	r, _, _ := util.EscapedString(s).Unquote()
	return r
}

func lc(s string) string {
	return strings.ToLower(s)
}

func uc(s string) string {
	return strings.ToUpper(s)
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func(interface{}, []Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Π<BoolExpr>	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolExpr : Π<BoolExpr1>	<< $0, nil >>`,
		Id:         "BoolExpr",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolExpr1 : Π<Val>	<< $0, nil >>`,
		Id:         "BoolExpr1",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolExpr1 : Π<BoolExpr> Λ<&> Π<BoolExpr1>	<< ast.NewBoolAndExpr($0, $2) >>`,
		Id:         "BoolExpr1",
		NTType:     2,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewBoolAndExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `BoolExpr1 : Π<BoolExpr> Λ<|> Π<BoolExpr1>	<< ast.NewBoolOrExpr($0, $2) >>`,
		Id:         "BoolExpr1",
		NTType:     2,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewBoolOrExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `BoolExpr1 : Λ<(> Π<BoolExpr> Λ<)>	<< ast.NewBoolGroupExpr($1) >>`,
		Id:         "BoolExpr1",
		NTType:     2,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewBoolGroupExpr(X[1])
		},
	},
	ProdTabEntry{
		String: `Val : Λ<true>	<< ast.TRUE, nil >>`,
		Id:         "Val",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.TRUE, nil
		},
	},
	ProdTabEntry{
		String: `Val : Λ<false>	<< ast.FALSE, nil >>`,
		Id:         "Val",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.FALSE, nil
		},
	},
	ProdTabEntry{
		String: `Val : Π<CompareExpr>	<< $0, nil >>`,
		Id:         "Val",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Val : Π<SubStringExpr>	<< $0, nil >>`,
		Id:         "Val",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CompareExpr : int_lit Λ<<> int_lit	<< ast.NewLessThanExpr($0, $2) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CompareExpr : int_lit Λ<>> int_lit	<< ast.NewLessThanExpr($2, $0) >>`,
		Id:         "CompareExpr",
		NTType:     4,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[2], X[0])
		},
	},
	ProdTabEntry{
		String: `SubStringExpr : string_lit Λ<in> string_lit	<< ast.NewSubStringExpr($0, $2) >>`,
		Id:         "SubStringExpr",
		NTType:     5,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewSubStringExpr(X[0], X[2])
		},
	},
}
