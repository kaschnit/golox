package ast

type Visitable interface {
	Accept(v AstVisitor) interface{}
}
type AstVisitor interface {
	VisitBinaryExpr(*BinaryExpr) interface{}
	VisitUnaryExpr(*UnaryExpr) interface{}
	VisitGroupingExpr(*GroupingExpr) interface{}
	VisitLiteralExpr(*LiteralExpr) interface{}
}
