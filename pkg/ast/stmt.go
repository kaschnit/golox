package ast

type Stmt interface {
	Accept(v AstVisitor) interface{}
}
