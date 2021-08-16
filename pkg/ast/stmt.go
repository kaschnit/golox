package ast

import "github.com/kaschnit/golox/pkg/token"

type Stmt interface {
	Accept(v AstVisitor) interface{}
}

type PrintStmt struct {
	Expression Expr
}

func (s *PrintStmt) Accept(v AstVisitor) interface{} {
	return v.VisitPrintStmt(s)
}

type ExprStmt struct {
	Expression Expr
}

func (s *ExprStmt) Accept(v AstVisitor) interface{} {
	return v.VisitExprStmt(s)
}

type IfStmt struct {
	Condition     Expr
	ThenStatement Stmt
	ElseStatement Stmt
}

func (s *IfStmt) Accept(v AstVisitor) interface{} {
	return v.VisitIfStmt(s)
}

type WhileStmt struct {
	Condition     Expr
	LoopStatement Stmt
}

func (s *WhileStmt) Accept(v AstVisitor) interface{} {
	return v.VisitWhileStmt(s)
}

type BlockStmt struct {
	Statements []Stmt
}

func (s *BlockStmt) Accept(v AstVisitor) interface{} {
	return v.VisitBlockStmt(s)
}

type VarStmt struct {
	Left  *token.Token
	Right Expr
}

func (s *VarStmt) Accept(v AstVisitor) interface{} {
	return v.VisitVarStmt(s)
}
