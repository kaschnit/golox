package ast

type Stmt interface {
	Accept(v AstVisitor) interface{}
}

type PrintStmt struct {
	Expr Expr
}

func (s *PrintStmt) Accept(v AstVisitor) interface{} {
	// TODO implement
	return nil
}

type ExprStmt struct {
	Expr Expr
}

func (s *ExprStmt) Accept(v AstVisitor) interface{} {
	// TODO implement
	return nil
}
