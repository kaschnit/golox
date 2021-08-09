package ast

import (
	"github.com/kaschnit/golox/pkg/scanner/token"
)

type Expr interface {
	Accept(v AstVisitor)
}

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (e *BinaryExpr) Accept(v AstVisitor) {
	v.VisitBinaryExpr(e)
}

type UnaryExpr struct {
	Operator token.Token
	Right    Expr
}

func (e *UnaryExpr) Accept(v AstVisitor) {
	v.VisitUnaryExpr(e)
}

type GroupingExpr struct {
	Expression Expr
}

func (e *GroupingExpr) Accept(v AstVisitor) {
	v.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	Value interface{}
}

func (e *LiteralExpr) Accept(v AstVisitor) {
	v.VisitLiteralExpr(e)
}
