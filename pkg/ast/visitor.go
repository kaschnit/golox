package ast

type AstVisitor interface {
	VisitBinaryExpr(*BinaryExpr)
	VisitUnaryExpr(*UnaryExpr)
	VisitGroupingExpr(*GroupingExpr)
	VisitLiteralExpr(*LiteralExpr)
}
