package emitter

import (
	"github.com/kaschnit/golox/pkg/ast"
)

// Implementation of AstVisitor that emits bytecode corresponding to the visited AST.
type AstEmitter struct{}

// Create an AstEmitter.
func NewAstEmitter() *AstEmitter {
	return &AstEmitter{}
}

func (p *AstEmitter) VisitAssignExpr(e *ast.AssignExpr) interface{} {
	return nil
}

func (ae *AstEmitter) VisitBinaryExpr(ex *ast.BinaryExpr) interface{} {
	return nil
}

func (ae *AstEmitter) VisitUnaryExpr(ex *ast.UnaryExpr) interface{} {

	return nil
}

func (ae *AstEmitter) VisitGroupingExpr(ex *ast.GroupingExpr) interface{} {
	return nil
}

func (ae *AstEmitter) VisitLiteralExpr(ex *ast.LiteralExpr) interface{} {
	return nil
}

func (ae *AstEmitter) VisitVarExpr(ex *ast.VarExpr) interface{} {
	return nil
}
