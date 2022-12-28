package ast

// An object that can visit the provided visitable AST nodes.
// Implements the "visitor" side of the visitor pattern.
type AstVisitor interface {
	VisitProgram(*Program) (interface{}, error)
	VisitPrintStmt(*PrintStmt) (interface{}, error)
	VisitReturnStmt(*ReturnStmt) (interface{}, error)
	VisitExprStmt(*ExprStmt) (interface{}, error)
	VisitIfStmt(*IfStmt) (interface{}, error)
	VisitWhileStmt(*WhileStmt) (interface{}, error)
	VisitAssignExpr(*AssignExpr) (interface{}, error)
	VisitCallExpr(*CallExpr) (interface{}, error)
	VisitBlockStmt(*BlockStmt) (interface{}, error)
	VisitClassStmt(*ClassStmt) (interface{}, error)
	VisitFunctionStmt(*FunctionStmt) (interface{}, error)
	VisitVarStmt(*VarStmt) (interface{}, error)
	VisitBinaryExpr(*BinaryExpr) (interface{}, error)
	VisitUnaryExpr(*UnaryExpr) (interface{}, error)
	VisitGroupingExpr(*GroupingExpr) (interface{}, error)
	VisitLiteralExpr(*LiteralExpr) (interface{}, error)
	VisitVarExpr(*VarExpr) (interface{}, error)
	VisitGetPropertyExpr(*GetPropertyExpr) (interface{}, error)
	VisitSetPropertyExpr(*SetPropertyExpr) (interface{}, error)
	VisitThisExpr(*ThisExpr) (interface{}, error)
}
