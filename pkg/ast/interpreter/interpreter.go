package interpreter

import (
	"github.com/kaschnit/golox/pkg/ast"
)

type AstInterpreter struct{}

func NewAstInterpreter() *AstInterpreter {
	return &AstInterpreter{}
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
