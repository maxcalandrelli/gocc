package ast

import (
	"github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er/iface"
)

type (
	StmtList []interface{}
	Stmt     string
)

func NewStmtList(stmt interface{}) (StmtList, error) {
	return StmtList{stmt}, nil
}

func AppendStmt(stmtList, stmt interface{}) (StmtList, error) {
	return append(stmtList.(StmtList), stmt), nil
}

func NewStmt(stmtList interface{}) (Stmt, error) {
	return Stmt(stmtList.(*iface.Token).Lit), nil
}
