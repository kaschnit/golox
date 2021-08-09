package printer

import (
	"fmt"
	"io"

	"github.com/kaschnit/golox/pkg/ast"
)

type AstPrinter struct {
	writer *io.Writer
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (a *AstPrinter) VisitBinaryExpr(e *ast.BinaryExpr) {
	fmt.Print(fmt.Sprintf("(%s ", e.Operator.Lexeme))
	e.Left.Accept(a)
	fmt.Print(" ")
	e.Right.Accept(a)
	fmt.Print(")")
}

func (a *AstPrinter) VisitUnaryExpr(e *ast.UnaryExpr) {
	fmt.Print(fmt.Sprintf("(%s ", e.Operator.Lexeme))
	e.Right.Accept(a)
	fmt.Print(")")
}

func (a *AstPrinter) VisitGroupingExpr(e *ast.GroupingExpr) {
	fmt.Print("(group ")
	e.Expression.Accept(a)
	fmt.Print(")")
}

func (a *AstPrinter) VisitLiteralExpr(e *ast.LiteralExpr) {
	if e.Value == nil {
		fmt.Print("nil")
	} else {
		fmt.Print(e.Value)
	}
}
