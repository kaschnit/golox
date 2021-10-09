package emitter

import (
	"github.com/kaschnit/golox/pkg/ast"
)

type AstEmitter struct{}

func NewAstEmitter() *AstEmitter {
	return &AstEmitter{}
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
