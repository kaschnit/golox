package analyzer

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/interpreter"
	loxerr "github.com/kaschnit/golox/pkg/errors"
)

type Scope map[string]bool

type ClassType int

const (
	ClassTypeNone ClassType = iota
	ClassTypeClass
)

type AstAnalyzer struct {
	interpreter        *interpreter.AstInterpreter
	scopes             []Scope
	resolutionDistance map[ast.Expr]int
	currentClassType   ClassType
}

func NewAstAnalyzer(interpreter *interpreter.AstInterpreter) *AstAnalyzer {
	return &AstAnalyzer{
		interpreter:        interpreter,
		scopes:             make([]Scope, 0),
		resolutionDistance: make(map[ast.Expr]int),
		currentClassType:   ClassTypeNone,
	}
}

func (r *AstAnalyzer) VisitProgram(prg *ast.Program) (interface{}, error) {
	for _, stmt := range prg.Statements {
		stmt.Accept(r)
	}
	return nil, nil
}

func (r *AstAnalyzer) VisitPrintStmt(s *ast.PrintStmt) (interface{}, error) {
	s.Expression.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitReturnStmt(s *ast.ReturnStmt) (interface{}, error) {
	if s.Expression != nil {
		s.Expression.Accept(r)
	}
	return nil, nil
}

func (r *AstAnalyzer) VisitExprStmt(s *ast.ExprStmt) (interface{}, error) {
	s.Expression.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitIfStmt(s *ast.IfStmt) (interface{}, error) {
	s.Condition.Accept(r)
	s.ThenStatement.Accept(r)
	if s.ElseStatement != nil {
		s.ElseStatement.Accept(r)
	}
	return nil, nil
}

func (r *AstAnalyzer) VisitWhileStmt(s *ast.WhileStmt) (interface{}, error) {
	s.Condition.Accept(r)
	s.LoopStatement.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitBlockStmt(s *ast.BlockStmt) (interface{}, error) {
	r.beginScope()
	for _, stmt := range s.Statements {
		stmt.Accept(r)
	}
	r.endScope()
	return nil, nil
}

func (r *AstAnalyzer) VisitClassStmt(s *ast.ClassStmt) (interface{}, error) {
	enclosingClassType := r.currentClassType
	defer func() {
		r.currentClassType = enclosingClassType
	}()

	r.currentClassType = ClassTypeClass

	r.defineName(s.Name.Lexeme)

	r.beginScope()

	r.defineName("this")

	for _, method := range s.Methods {
		r.resolveFunction(method)
	}

	r.endScope()
	return nil, nil
}

func (r *AstAnalyzer) VisitFunctionStmt(s *ast.FunctionStmt) (interface{}, error) {
	r.defineName(s.Name.Lexeme)
	r.resolveFunction(s)
	return nil, nil
}

func (r *AstAnalyzer) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	r.declareName(s.Left.Lexeme)
	if s.Right != nil {
		s.Right.Accept(r)
	}
	r.defineName(s.Left.Lexeme)
	return nil, nil
}

func (r *AstAnalyzer) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	e.Right.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	e.Callee.Accept(r)
	for _, arg := range e.Args {
		arg.Accept(r)
	}
	return nil, nil
}

func (r *AstAnalyzer) VisitBinaryExpr(e *ast.BinaryExpr) (interface{}, error) {
	e.Left.Accept(r)
	e.Right.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitUnaryExpr(e *ast.UnaryExpr) (interface{}, error) {
	e.Right.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	e.Expression.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	return nil, nil
}

func (r *AstAnalyzer) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	if len(r.scopes) > 0 {
		if val, ok := r.scopes[len(r.scopes)-1][e.Name.Lexeme]; ok && !val {
			loxerr.AtToken(e.Name, "Can't read local variable in its own initializer.")
		}
	}
	return nil, nil
}

func (r *AstAnalyzer) VisitGetPropertyExpr(e *ast.GetPropertyExpr) (interface{}, error) {
	e.ParentObject.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitSetPropertyExpr(e *ast.SetPropertyExpr) (interface{}, error) {
	e.Value.Accept(r)
	e.ParentObject.Accept(r)
	return nil, nil
}

func (r *AstAnalyzer) VisitThisExpr(e *ast.ThisExpr) (interface{}, error) {
	if r.currentClassType == ClassTypeNone {
		return nil, loxerr.AtToken(e.Keyword, "Can't use 'this' outside of a class.")
	}
	return nil, nil
}

func (r *AstAnalyzer) resolveFunction(f *ast.FunctionStmt) {
	r.beginScope()
	for _, param := range f.Params {
		r.defineName(param.Lexeme)
	}
	for _, stmt := range f.Body {
		stmt.Accept(r)
	}
	r.endScope()
}

func (r *AstAnalyzer) beginScope() *Scope {
	newScope := make(Scope)
	r.scopes = append(r.scopes, newScope)
	return &newScope
}

func (r *AstAnalyzer) endScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

func (r *AstAnalyzer) declareName(name string) {
	if len(r.scopes) == 0 {
		return
	}
	r.scopes[len(r.scopes)-1][name] = false
}

func (r *AstAnalyzer) defineName(name string) {
	if len(r.scopes) == 0 {
		return
	}
	r.scopes[len(r.scopes)-1][name] = true
}
