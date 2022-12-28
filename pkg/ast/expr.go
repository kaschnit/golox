package ast

import (
	"github.com/kaschnit/golox/pkg/token"
)

// Represents any expression AST node.
type Expr interface {
	Accept(v AstVisitor) (interface{}, error)
}

// Represents an assignment expression AST node.
type AssignExpr struct {
	Left  *token.Token
	Right Expr
}

func (e *AssignExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitAssignExpr(e)
}

// Represents a call expression AST node.
type CallExpr struct {
	Callee    Expr
	OpenParen *token.Token
	Args      []Expr
}

func (e *CallExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitCallExpr(e)
}

// Represents a binary expression AST node.
type BinaryExpr struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (e *BinaryExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitBinaryExpr(e)
}

// Represents a unary expression AST node.
type UnaryExpr struct {
	Operator *token.Token
	Right    Expr
}

func (e *UnaryExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitUnaryExpr(e)
}

// Represents a grouping expression AST node.
type GroupingExpr struct {
	Expression Expr
}

func (e *GroupingExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitGroupingExpr(e)
}

// Represents a literal expression AST node.
type LiteralExpr struct {
	Value interface{}
}

func (e *LiteralExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitLiteralExpr(e)
}

// Represents a variable usage expression AST node.
type VarExpr struct {
	Name *token.Token
}

func (e *VarExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitVarExpr(e)
}

type GetPropertyExpr struct {
	Name         *token.Token
	ParentObject Expr
}

func (e *GetPropertyExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitGetPropertyExpr(e)
}

type SetPropertyExpr struct {
	Name         *token.Token
	Value        Expr
	ParentObject Expr
}

func (e *SetPropertyExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitSetPropertyExpr(e)
}

type ThisExpr struct {
	Keyword *token.Token
}

func (e *ThisExpr) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitThisExpr(e)
}
