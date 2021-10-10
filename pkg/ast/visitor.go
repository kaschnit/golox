package ast

type Visitable interface {
	Accept(v AstVisitor) interface{}
}
type AstVisitor interface {
	VisitProgram(*Program) interface{}
	VisitPrintStmt(*PrintStmt) interface{}
	VisitExprStmt(*ExprStmt) interface{}
	VisitIfStmt(*IfStmt) interface{}
	VisitWhileStmt(*WhileStmt) interface{}
	VisitBlockStmt(*BlockStmt) interface{}
	VisitVarStmt(*VarStmt) interface{}
	VisitBinaryExpr(*BinaryExpr) interface{}
	VisitUnaryExpr(*UnaryExpr) interface{}
	VisitGroupingExpr(*GroupingExpr) interface{}
	VisitLiteralExpr(*LiteralExpr) interface{}
	VisitVarExpr(*VarExpr) interface{}
}
