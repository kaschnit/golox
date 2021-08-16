package printer

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
)

type AstPrinter struct {
	indent int
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{indent: 0}
}

func (p *AstPrinter) VisitProgram(prg *ast.Program) interface{} {
	for i := 0; i < len(prg.Statements); i++ {
		prg.Statements[i].Accept(p)
	}
	return nil
}

func (p *AstPrinter) VisitPrintStmt(s *ast.PrintStmt) interface{} {
	p.printTabbing()
	fmt.Print("(print ")
	s.Expression.Accept(p)
	fmt.Println(");")
	return nil
}

func (p *AstPrinter) VisitExprStmt(s *ast.ExprStmt) interface{} {
	p.printTabbing()
	fmt.Print("(")
	s.Expression.Accept(p)
	fmt.Println(");")
	return nil
}

func (p *AstPrinter) VisitIfStmt(s *ast.IfStmt) interface{} {
	p.printTabbing()
	fmt.Print("if (condition ")
	s.Condition.Accept(p)
	fmt.Println("):")
	p.indent++
	s.ThenStatement.Accept(p)
	p.indent--
	if s.ElseStatement != nil {
		p.indent++
		s.ElseStatement.Accept(p)
		p.indent--
	}
	return nil
}

func (p *AstPrinter) VisitWhileStmt(s *ast.WhileStmt) interface{} {
	p.printTabbing()
	fmt.Print("while (condition ")
	s.Condition.Accept(p)
	fmt.Println("):")
	p.indent++
	s.LoopStatement.Accept(p)
	p.indent--
	return nil
}

func (p *AstPrinter) VisitBlockStmt(s *ast.BlockStmt) interface{} {
	p.printTabbing()
	fmt.Println("{")
	p.indent++
	for i := 0; i < len(s.Statements); i++ {
		s.Statements[i].Accept(p)
	}
	p.indent--
	p.printTabbing()
	fmt.Println("}")
	return nil
}

func (p *AstPrinter) VisitVarStmt(s *ast.VarStmt) interface{} {
	p.printTabbing()
	fmt.Print("(var ")
	fmt.Println(");")
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
	} else if strVal, ok := e.Value.(string); ok {
		fmt.Printf(`"%s"`, strVal)
	} else {
		fmt.Print(e.Value)
	}
	return nil
}

func (p *AstPrinter) printTabbing() {
	for i := 0; i < p.indent; i++ {
		fmt.Print("  ")
	}
}
