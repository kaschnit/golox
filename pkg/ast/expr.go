package ast

import (
	"github.com/kaschnit/golox/pkg/token"
)

type Expr interface {
	Accept(v AstVisitor) interface{}
}

type BinaryExpr struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (e *BinaryExpr) Accept(v AstVisitor) interface{} {
	return v.VisitBinaryExpr(e)
}

type UnaryExpr struct {
	Operator *token.Token
	Right    Expr
}

func (e *UnaryExpr) Accept(v AstVisitor) interface{} {
	return v.VisitUnaryExpr(e)
}

type GroupingExpr struct {
	Expression Expr
}

func (e *GroupingExpr) Accept(v AstVisitor) interface{} {
	return v.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	Value interface{}
}

func (e *LiteralExpr) Accept(v AstVisitor) interface{} {
	return v.VisitLiteralExpr(e)
}

type VarExpr struct {
	Name string
}

func (e *VarExpr) Accept(v AstVisitor) interface{} {
	return v.VisitVarExpr(e)
}
