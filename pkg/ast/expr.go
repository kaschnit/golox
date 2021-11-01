package ast

import (
	"github.com/kaschnit/golox/pkg/token"
)

type Expr interface {
	Accept(v AstVisitor) (interface{}, error)
}

type AssignExpr struct {
	Left  *token.Token
	Right Expr
}

func (e *AssignExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitAssignExpr(e)
}

type CallExpr struct {
	Callee    Expr
	OpenParen *token.Token
	Args      []Expr
}

func (e *CallExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitCallExpr(e)
}

type BinaryExpr struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (e *BinaryExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitBinaryExpr(e)
}

type UnaryExpr struct {
	Operator *token.Token
	Right    Expr
}

func (e *UnaryExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitUnaryExpr(e)
}

type GroupingExpr struct {
	Expression Expr
}

func (e *GroupingExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	Value interface{}
}

func (e *LiteralExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitLiteralExpr(e)
}

type VarExpr struct {
	Name *token.Token
}

func (e *VarExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitVarExpr(e)
}
