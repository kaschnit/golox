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

func (p *AstPrinter) VisitBinaryExpr(e *ast.BinaryExpr) interface{} {
	fmt.Print(fmt.Sprintf("(%s ", e.Operator.Lexeme))
	e.Left.Accept(p)
	fmt.Print(" ")
	e.Right.Accept(p)
	fmt.Print(")")
	return nil
}

func (p *AstPrinter) VisitUnaryExpr(e *ast.UnaryExpr) interface{} {
	fmt.Print(fmt.Sprintf("(%s ", e.Operator.Lexeme))
	e.Right.Accept(p)
	fmt.Print(")")
	return nil
}

func (p *AstPrinter) VisitGroupingExpr(e *ast.GroupingExpr) interface{} {
	fmt.Print("(group ")
	e.Expression.Accept(p)
	fmt.Print(")")
	return nil
}

func (p *AstPrinter) VisitLiteralExpr(e *ast.LiteralExpr) interface{} {
	if e.Value == nil {
		fmt.Print("nil")
	} else {
		fmt.Print(e.Value)
	}
	return nil
}
