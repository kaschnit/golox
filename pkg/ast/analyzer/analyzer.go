package analyzer

import (
	"github.com/hashicorp/go-multierror"
	"github.com/kaschnit/golox/pkg/ast"
	loxerr "github.com/kaschnit/golox/pkg/errors"
)

type Scope map[string]bool

type ClassType int

const (
	ClassTypeNone ClassType = iota
	ClassTypeClass
)

type FunctionType int

const (
	FunctionTypeNone FunctionType = iota
	FunctionTypeFunction
	FunctionTypeMethod
	FunctionTypeConstructor
)

type AstAnalyzer struct {
	scopes              []Scope
	resolutionDistance  map[ast.Expr]int
	currentClassType    ClassType
	currentFunctionType FunctionType
}

func NewAstAnalyzer() *AstAnalyzer {
	return &AstAnalyzer{
		scopes:              make([]Scope, 0),
		resolutionDistance:  make(map[ast.Expr]int),
		currentClassType:    ClassTypeNone,
		currentFunctionType: FunctionTypeNone,
	}
}

func (r *AstAnalyzer) VisitProgram(prg *ast.Program) (interface{}, error) {
	errs := new(multierror.Error)
	for _, stmt := range prg.Statements {
		_, err := stmt.Accept(r)
		errs = multierror.Append(errs, err)
	}
	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitPrintStmt(s *ast.PrintStmt) (interface{}, error) {
	_, err := s.Expression.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitReturnStmt(s *ast.ReturnStmt) (interface{}, error) {
	errs := new(multierror.Error)
	if s.Expression != nil {
		if r.currentFunctionType == FunctionTypeConstructor {
			err := loxerr.AtToken(s.Keyword, "Can't return a value from a constructor.")
			errs = multierror.Append(errs, err)
		}

		_, err := s.Expression.Accept(r)
		errs = multierror.Append(errs, err)
	}

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitExprStmt(s *ast.ExprStmt) (interface{}, error) {
	_, err := s.Expression.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitIfStmt(s *ast.IfStmt) (interface{}, error) {
	errs := new(multierror.Error)

	_, err := s.Condition.Accept(r)
	errs = multierror.Append(errs, err)

	_, err = s.ThenStatement.Accept(r)
	errs = multierror.Append(errs, err)

	if s.ElseStatement != nil {
		_, err = s.ElseStatement.Accept(r)
		errs = multierror.Append(errs, err)
	}
	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitWhileStmt(s *ast.WhileStmt) (interface{}, error) {
	errs := new(multierror.Error)

	_, err := s.Condition.Accept(r)
	errs = multierror.Append(errs, err)

	_, err = s.LoopStatement.Accept(r)
	errs = multierror.Append(errs, err)

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitBlockStmt(s *ast.BlockStmt) (interface{}, error) {
	errs := new(multierror.Error)

	r.beginScope()
	for _, stmt := range s.Statements {
		_, err := stmt.Accept(r)
		errs = multierror.Append(errs, err)
	}
	r.endScope()

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitClassStmt(s *ast.ClassStmt) (interface{}, error) {
	errs := new(multierror.Error)

	enclosingClassType := r.currentClassType
	defer func() {
		r.currentClassType = enclosingClassType
	}()

	r.currentClassType = ClassTypeClass

	r.defineName(s.Name.Lexeme)

	r.beginScope()

	r.defineName("this")

	if s.Constructor != nil {
		err := r.resolveFunction(s.Constructor, FunctionTypeConstructor)
		errs = multierror.Append(errs, err)
	}
	for _, method := range s.Methods {
		err := r.resolveFunction(method, FunctionTypeMethod)
		errs = multierror.Append(errs, err)
	}

	r.endScope()
	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitFunctionStmt(s *ast.FunctionStmt) (interface{}, error) {
	r.defineName(s.Name.Lexeme)
	err := r.resolveFunction(s, FunctionTypeFunction)
	return nil, err
}

func (r *AstAnalyzer) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	errs := new(multierror.Error)

	r.declareName(s.Left.Lexeme)
	if s.Right != nil {
		_, err := s.Right.Accept(r)
		errs = multierror.Append(errs, err)
	}
	r.defineName(s.Left.Lexeme)

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	_, err := e.Right.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	errs := new(multierror.Error)

	e.Callee.Accept(r)
	for _, arg := range e.Args {
		_, err := arg.Accept(r)
		errs = multierror.Append(errs, err)
	}
	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitBinaryExpr(e *ast.BinaryExpr) (interface{}, error) {
	errs := new(multierror.Error)

	_, err := e.Left.Accept(r)
	errs = multierror.Append(errs, err)

	_, err = e.Right.Accept(r)
	errs = multierror.Append(errs, err)

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitUnaryExpr(e *ast.UnaryExpr) (interface{}, error) {
	_, err := e.Right.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	_, err := e.Expression.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	return nil, nil
}

func (r *AstAnalyzer) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	errs := new(multierror.Error)

	if len(r.scopes) > 0 {
		if val, ok := r.scopes[len(r.scopes)-1][e.Name.Lexeme]; ok && !val {
			err := loxerr.AtToken(e.Name, "Can't read local variable in its own initializer.")
			errs = multierror.Append(errs, err)
		}
	}

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitGetPropertyExpr(e *ast.GetPropertyExpr) (interface{}, error) {
	_, err := e.ParentObject.Accept(r)
	return nil, err
}

func (r *AstAnalyzer) VisitSetPropertyExpr(e *ast.SetPropertyExpr) (interface{}, error) {
	errs := new(multierror.Error)

	_, err := e.Value.Accept(r)
	errs = multierror.Append(errs, err)

	_, err = e.ParentObject.Accept(r)
	errs = multierror.Append(errs, err)

	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) VisitThisExpr(e *ast.ThisExpr) (interface{}, error) {
	errs := new(multierror.Error)
	if r.currentClassType == ClassTypeNone {
		err := loxerr.AtToken(e.Keyword, "Can't use 'this' outside of a class.")
		errs = multierror.Append(errs, err)
	}
	return nil, errs.ErrorOrNil()
}

func (r *AstAnalyzer) resolveFunction(f *ast.FunctionStmt, kind FunctionType) error {
	errs := new(multierror.Error)

	enclosingFunctionType := r.currentFunctionType
	defer func() {
		r.currentFunctionType = enclosingFunctionType
	}()
	r.currentFunctionType = kind

	r.beginScope()
	for _, param := range f.Params {
		r.defineName(param.Lexeme)
	}
	for _, stmt := range f.Body {
		_, err := stmt.Accept(r)
		errs = multierror.Append(errs, err)
	}
	r.endScope()

	return errs.ErrorOrNil()
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
