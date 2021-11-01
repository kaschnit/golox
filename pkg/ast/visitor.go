package ast

type Visitable interface {
	Accept(v AstVisitor) (interface{}, error)
}
type AstVisitor interface {
	VisitProgram(*Program) (interface{}, error)
	VisitPrintStmt(*PrintStmt) (interface{}, error)
	VisitExprStmt(*ExprStmt) (interface{}, error)
	VisitIfStmt(*IfStmt) (interface{}, error)
	VisitWhileStmt(*WhileStmt) (interface{}, error)
	VisitAssignExpr(*AssignExpr) (interface{}, error)
	VisitCallExpr(*CallExpr) (interface{}, error)
	VisitBlockStmt(*BlockStmt) (interface{}, error)
	VisitVarStmt(*VarStmt) (interface{}, error)
	VisitBinaryExpr(*BinaryExpr) (interface{}, error)
	VisitUnaryExpr(*UnaryExpr) (interface{}, error)
	VisitGroupingExpr(*GroupingExpr) (interface{}, error)
	VisitLiteralExpr(*LiteralExpr) (interface{}, error)
	VisitVarExpr(*VarExpr) (interface{}, error)
}
