package ast

import "github.com/kaschnit/golox/pkg/token"

// Represents any statement AST node.
type Stmt interface {
	Accept(v AstVisitor) (interface{}, error)
}

// Represents a print statement AST node.
type PrintStmt struct {
	Expression Expr
}

func (s *PrintStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitPrintStmt(s)
}

// Represents a return statement AS node.
type ReturnStmt struct {
	Keyword    *token.Token
	Expression Expr
}

func (s *ReturnStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitReturnStmt(s)
}

// Represents an expression statement AST node.
type ExprStmt struct {
	Expression Expr
}

func (s *ExprStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitExprStmt(s)
}

// Represents an if statement AST node.
type IfStmt struct {
	Condition     Expr
	ThenStatement Stmt
	ElseStatement Stmt
}

func (s *IfStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitIfStmt(s)
}

// Represents a while loop AST node.
type WhileStmt struct {
	Condition     Expr
	LoopStatement Stmt
}

func (s *WhileStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitWhileStmt(s)
}

// Represents a block AST node.
type BlockStmt struct {
	Statements []Stmt
}

func (s *BlockStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitBlockStmt(s)
}

// Represents a class declaration statement AST node.
type ClassStmt struct {
	Name          *token.Token
	Constructor   *FunctionStmt
	Methods       []*FunctionStmt
	StaticMethods []*FunctionStmt
}

func (s *ClassStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitClassStmt(s)
}

// Represents a function declaration statement AST node.
type FunctionStmt struct {
	Name   *token.Token
	Params []*token.Token
	Body   []Stmt
}

func (s *FunctionStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitFunctionStmt(s)
}

// Represents a var declaration statement AST node.
type VarStmt struct {
	Left  *token.Token
	Right Expr
}

func (s *VarStmt) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitVarStmt(s)
}
