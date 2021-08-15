package printer

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
)

type AstPrinter struct{}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) VisitProgram(prg *ast.Program) interface{} {
	for i := 0; i < len(prg.Statements); i++ {
		prg.Statements[i].Accept(p)
	}
	return nil
}

func (p *AstPrinter) VisitPrintStmt(ps *ast.PrintStmt) interface{} {
	fmt.Print("print ")
	ps.Expr.Accept(p)
	fmt.Println(";")
	return nil
}

func (p *AstPrinter) VisitExprStmt(ps *ast.ExprStmt) interface{} {
	ps.Expr.Accept(p)
	fmt.Println(";")
	return nil
}

func (p *AstPrinter) VisitBinaryExpr(e *ast.BinaryExpr) interface{} {
	fmt.Printf("(%s ", e.Operator.Lexeme)
	e.Left.Accept(p)
	fmt.Print(" ")
	e.Right.Accept(p)
	fmt.Print(")")
	return nil
}

func (p *AstPrinter) VisitUnaryExpr(e *ast.UnaryExpr) interface{} {
	fmt.Printf("(%s ", e.Operator.Lexeme)
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
