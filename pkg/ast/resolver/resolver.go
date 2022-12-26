package resolver

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/interpreter"
	loxerr "github.com/kaschnit/golox/pkg/errors"
)

type Scope map[string]bool

type AstResolver struct {
	interpreter        *interpreter.AstInterpreter
	scopes             []Scope
	resolutionDistance map[ast.Expr]int
}

func NewAstResolver(interpreter *interpreter.AstInterpreter) *AstResolver {
	return &AstResolver{
		interpreter:        interpreter,
		scopes:             make([]Scope, 0),
		resolutionDistance: make(map[ast.Expr]int),
	}
}

func (r *AstResolver) VisitProgram(prg *ast.Program) (interface{}, error) {
	for _, stmt := range prg.Statements {
		stmt.Accept(r)
	}
	return nil, nil
}

func (r *AstResolver) VisitPrintStmt(s *ast.PrintStmt) (interface{}, error) {
	s.Expression.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitReturnStmt(s *ast.ReturnStmt) (interface{}, error) {
	if s.Expression != nil {
		s.Expression.Accept(r)
	}
	return nil, nil
}

func (r *AstResolver) VisitExprStmt(s *ast.ExprStmt) (interface{}, error) {
	s.Expression.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitIfStmt(s *ast.IfStmt) (interface{}, error) {
	s.Condition.Accept(r)
	s.ThenStatement.Accept(r)
	if s.ElseStatement != nil {
		s.ElseStatement.Accept(r)
	}
	return nil, nil
}

func (r *AstResolver) VisitWhileStmt(s *ast.WhileStmt) (interface{}, error) {
	s.Condition.Accept(r)
	s.LoopStatement.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitBlockStmt(s *ast.BlockStmt) (interface{}, error) {
	r.beginScope()
	for _, stmt := range s.Statements {
		stmt.Accept(r)
	}
	r.endScope()
	return nil, nil
}

func (r *AstResolver) VisitFunctionStmt(s *ast.FunctionStmt) (interface{}, error) {
	r.declareName(s.Name.Lexeme)
	r.defineName(s.Name.Lexeme)
	r.resolveFunction(s)
	return nil, nil
}

func (r *AstResolver) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	r.declareName(s.Left.Lexeme)
	if s.Right != nil {
		s.Right.Accept(r)
	}
	r.defineName(s.Left.Lexeme)
	return nil, nil
}

func (r *AstResolver) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	e.Right.Accept(r)
	r.resolveLocal(e, e.Left.Lexeme)
	return nil, nil
}

func (r *AstResolver) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	e.Callee.Accept(r)
	for _, arg := range e.Args {
		arg.Accept(r)
	}
	return nil, nil
}

func (r *AstResolver) VisitBinaryExpr(e *ast.BinaryExpr) (interface{}, error) {
	e.Left.Accept(r)
	e.Right.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitUnaryExpr(e *ast.UnaryExpr) (interface{}, error) {
	e.Right.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	e.Expression.Accept(r)
	return nil, nil
}

func (r *AstResolver) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	return nil, nil
}

func (r *AstResolver) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	if len(r.scopes) > 0 {
		if val, ok := r.scopes[len(r.scopes)-1][e.Name.Lexeme]; ok && !val {
			loxerr.AtToken(e.Name, "Can't read local variable in its own initializer.")
		}
	}
	r.resolveLocal(e, e.Name.Lexeme)
	return nil, nil
}

func (r *AstResolver) resolveLocal(expr ast.Expr, name string) {
	for i := len(r.scopes) - 1; i >= 0; i-- {
		if _, ok := r.scopes[i][name]; ok {
			r.interpreter.Resolve(expr, len(r.scopes)-1-i)
		}
	}
}

func (r *AstResolver) resolveFunction(f *ast.FunctionStmt) {
	r.beginScope()
	for _, param := range f.Params {
		r.declareName(param.Lexeme)
		r.defineName(param.Lexeme)
	}
	f.Body.Accept(r)
	r.endScope()
}

func (r *AstResolver) beginScope() {
	r.scopes = append(r.scopes, make(Scope))
}

func (r *AstResolver) endScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

func (r *AstResolver) declareName(name string) {
	if len(r.scopes) == 0 {
		return
	}
	r.scopes[len(r.scopes)-1][name] = false
}

func (r *AstResolver) defineName(name string) {
	if len(r.scopes) == 0 {
		return
	}
	r.scopes[len(r.scopes)-1][name] = true
}
