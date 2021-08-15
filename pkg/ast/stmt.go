package ast

type Stmt interface {
	Accept(v AstVisitor) interface{}
}

type PrintStmt struct {
	Expr Expr
}

func (s *PrintStmt) Accept(v AstVisitor) interface{} {
	return v.VisitPrintStmt(s)
}

type ExprStmt struct {
	Expr Expr
}

func (s *ExprStmt) Accept(v AstVisitor) interface{} {
	return v.VisitExprStmt(s)
}
