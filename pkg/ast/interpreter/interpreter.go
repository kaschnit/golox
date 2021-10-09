package interpreter

import (
	"github.com/kaschnit/golox/pkg/ast"
)

type AstInterpreter struct{}

func NewAstInterpreter() *AstInterpreter {
	return &AstInterpreter{}
}

func (i *AstInterpreter) VisitProgram(*ast.Program) interface{} {
	return nil
}

func (i *AstInterpreter) VisitPrintStmt(*ast.PrintStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitExprStmt(*ast.ExprStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitIfStmt(*ast.IfStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitWhileStmt(*ast.WhileStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitBlockStmt(*ast.BlockStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitVarStmt(*ast.VarStmt) interface{} {
	return nil
}

func (i *AstInterpreter) VisitBinaryExpr(e *ast.BinaryExpr) interface{} {
	return nil
}

func (i *AstInterpreter) VisitUnaryExpr(e *ast.UnaryExpr) interface{} {
	return nil
}

func (i *AstInterpreter) VisitGroupingExpr(e *ast.GroupingExpr) interface{} {
	return nil
}

func (i *AstInterpreter) VisitLiteralExpr(e *ast.LiteralExpr) interface{} {
	return nil
}
