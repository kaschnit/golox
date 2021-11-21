package printer

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
)

// Implementation of AstVisitor that prints the visited AST.
type AstPrinter struct {
	indent int
}

// Create an AstPrinter.
func NewAstPrinter() *AstPrinter {
	return &AstPrinter{indent: 0}
}

func (p *AstPrinter) VisitProgram(prg *ast.Program) (interface{}, error) {
	for i := 0; i < len(prg.Statements); i++ {
		prg.Statements[i].Accept(p)
	}
	return nil, nil
}

func (p *AstPrinter) VisitPrintStmt(s *ast.PrintStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Print("(print ")
	s.Expression.Accept(p)
	fmt.Println(");")
	return nil, nil
}

func (p *AstPrinter) VisitReturnStmt(s *ast.ReturnStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Print("(return ")
	s.Expression.Accept(p)
	fmt.Println(");")
	return nil, nil
}

func (p *AstPrinter) VisitExprStmt(s *ast.ExprStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Print("(")
	s.Expression.Accept(p)
	fmt.Println(");")
	return nil, nil
}

func (p *AstPrinter) VisitIfStmt(s *ast.IfStmt) (interface{}, error) {
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
	return nil, nil
}

func (p *AstPrinter) VisitWhileStmt(s *ast.WhileStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Print("while (condition ")
	s.Condition.Accept(p)
	fmt.Println("):")
	p.indent++
	s.LoopStatement.Accept(p)
	p.indent--
	return nil, nil
}

func (p *AstPrinter) VisitBlockStmt(s *ast.BlockStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Println("{")
	p.indent++
	for i := 0; i < len(s.Statements); i++ {
		s.Statements[i].Accept(p)
	}
	p.indent--
	p.printTabbing()
	fmt.Println("}")
	return nil, nil
}

func (p *AstPrinter) VisitFunctionStmt(s *ast.FunctionStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Println("(func)")
	return nil, nil
}

func (p *AstPrinter) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	p.printTabbing()
	fmt.Printf("(var %s = ", s.Left.Lexeme)
	if s.Right != nil {
		s.Right.Accept(p)
	} else {
		fmt.Print("nil")
	}
	fmt.Println(");")
	return nil, nil
}

func (p *AstPrinter) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	fmt.Printf("(assign %s ", e.Left.Lexeme)
	e.Right.Accept(p)
	fmt.Print(")")
	return nil, nil
}

func (p *AstPrinter) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	fmt.Print("(call ")
	e.Callee.Accept(p)
	fmt.Print("(")
	for i, v := range e.Args {
		v.Accept(p)
		if i != len(e.Args)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("))")
	return nil, nil
}

func (p *AstPrinter) VisitBinaryExpr(e *ast.BinaryExpr) (interface{}, error) {
	fmt.Printf("(%s ", e.Operator.Lexeme)
	e.Left.Accept(p)
	fmt.Print(" ")
	e.Right.Accept(p)
	fmt.Print(")")
	return nil, nil
}

func (p *AstPrinter) VisitUnaryExpr(e *ast.UnaryExpr) (interface{}, error) {
	fmt.Printf("(%s ", e.Operator.Lexeme)
	e.Right.Accept(p)
	fmt.Print(")")
	return nil, nil
}

func (p *AstPrinter) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	fmt.Print("(group ")
	e.Expression.Accept(p)
	fmt.Print(")")
	return nil, nil
}

func (p *AstPrinter) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	if e.Value == nil {
		fmt.Print("nil")
	} else if strVal, ok := e.Value.(string); ok {
		fmt.Printf(`"%s"`, strVal)
	} else {
		fmt.Print(e.Value)
	}
	return nil, nil
}

func (p *AstPrinter) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	fmt.Printf("(var %s)", e.Name.Lexeme)
	return nil, nil
}

func (p *AstPrinter) printTabbing() {
	for i := 0; i < p.indent; i++ {
		fmt.Print("  ")
	}
}
