package interpreter

import (
	"fmt"
	"time"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/conversion"
	"github.com/kaschnit/golox/pkg/environment"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

// Implementation of AstVisitor that interprets the visited AST directly
type AstInterpreter struct {
	env             *environment.Environment
	globals         *environment.Environment
	localResolution map[ast.Expr]int
}

// Create an AstInterpreter.
func NewAstInterpreter() *AstInterpreter {
	globals := environment.NewEnvironmentWithVars(nil, map[string]interface{}{
		"clock": NewNativeFunction(
			0,
			func(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
				return time.Now().Unix(), nil
			},
		),
	})
	return &AstInterpreter{
		env:             globals,
		globals:         globals,
		localResolution: make(map[ast.Expr]int),
	}
}

func (a *AstInterpreter) VisitProgram(p *ast.Program) (interface{}, error) {
	for i := 0; i < len(p.Statements); i++ {
		_, err := p.Statements[i].Accept(a)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (a *AstInterpreter) VisitPrintStmt(s *ast.PrintStmt) (interface{}, error) {
	value, err := s.Expression.Accept(a)
	if err != nil {
		return nil, err
	}

	fmt.Print(value)
	return nil, nil
}

func (a *AstInterpreter) VisitReturnStmt(s *ast.ReturnStmt) (interface{}, error) {
	value, err := s.Expression.Accept(a)
	if err != nil {
		return nil, err
	}
	return nil, NewReturn(value)
}

func (a *AstInterpreter) VisitExprStmt(s *ast.ExprStmt) (interface{}, error) {
	return s.Expression.Accept(a)
}

func (a *AstInterpreter) VisitIfStmt(s *ast.IfStmt) (interface{}, error) {
	cond, err := s.Condition.Accept(a)
	if err != nil {
		return nil, err
	}

	if conversion.IsTruthy(cond) {
		_, err = s.ThenStatement.Accept(a)
	} else if s.ElseStatement != nil {
		_, err = s.ElseStatement.Accept(a)
	}
	return nil, err
}

func (a *AstInterpreter) VisitWhileStmt(s *ast.WhileStmt) (interface{}, error) {
	cond, err := s.Condition.Accept(a)
	if err != nil {
		return nil, err
	}

	for conversion.IsTruthy(cond) {
		s.LoopStatement.Accept(a)

		cond, err = s.Condition.Accept(a)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (a *AstInterpreter) VisitBlockStmt(s *ast.BlockStmt) (interface{}, error) {
	// Execute the block in a new child environment.
	return nil, a.ExecuteBlock(s.Statements, a.env.NewChild())
}

func (a *AstInterpreter) VisitFunctionStmt(s *ast.FunctionStmt) (interface{}, error) {
	function := NewLoxFunction(s, a.env)
	_, exists := a.env.Get(s.Name.Lexeme)
	if exists {
		return nil, loxerr.Runtime(s.Name, fmt.Sprintf("Name '%s' already defined", s.Name.Lexeme))
	}
	a.env.Set(s.Name.Lexeme, function)
	return nil, nil
}

func (a *AstInterpreter) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	_, exists := a.env.Get(s.Left.Lexeme)
	if exists {
		return nil, loxerr.Runtime(s.Left, fmt.Sprintf("Name '%s' already defined", s.Left.Lexeme))
	}

	if s.Right == nil {
		a.env.Set(s.Left.Lexeme, nil)
	} else {
		value, err := s.Right.Accept(a)
		if err != nil {
			return nil, err
		}
		a.env.Set(s.Left.Lexeme, value)
	}

	return nil, nil
}

func (a *AstInterpreter) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	value, err := e.Right.Accept(a)
	if err != nil {
		return nil, err
	}

	if distance, ok := a.localResolution[e]; ok {
		a.env.SetAt(e.Left.Lexeme, distance, value)
		return value, nil
	}

	exists := a.globals.Replace(e.Left.Lexeme, value)
	if !exists {
		return nil, loxerr.Runtime(e.Left, fmt.Sprintf("Variable '%s' not defined", e.Left.Lexeme))
	}

	return value, nil
}

func (a *AstInterpreter) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	callee, err := e.Callee.Accept(a)
	if err != nil {
		return nil, err
	}

	callable, ok := callee.(Callable)
	if !ok {
		return nil, loxerr.Runtime(e.OpenParen,
			fmt.Sprintf("Expression '%v' is not callable", callee))
	}

	if len(e.Args) != callable.Arity() {
		return nil, loxerr.Runtime(e.OpenParen,
			fmt.Sprintf("Expected %d args, got %d.", callable.Arity(), len(e.Args)))
	}

	argList := make([]interface{}, 0)
	for _, argExpr := range e.Args {
		argValue, err := argExpr.Accept(a)
		if err != nil {
			return nil, err
		}
		argList = append(argList, argValue)
	}

	result, err := callable.Call(a, argList)
	return result, err
}

func (a *AstInterpreter) VisitBinaryExpr(e *ast.BinaryExpr) (interface{}, error) {
	lhs, err := e.Left.Accept(a)
	if err != nil {
		return nil, err
	}
	lhsFloat, isLhsFloat := conversion.ToFloat(lhs)

	rhs, err := e.Right.Accept(a)
	if err != nil {
		return nil, err
	}
	rhsFloat, isRhsFloat := conversion.ToFloat(rhs)

	invalidOperatorMsg := fmt.Sprintf("Invalid operator '%s'", e.Operator.Lexeme)

	switch e.Operator.Type {
	case tokentype.MINUS:
		if isLhsFloat && isRhsFloat {
			return lhsFloat - rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.PLUS:
		if isLhsFloat && isRhsFloat {
			return lhsFloat + rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.SLASH:
		if isLhsFloat && isRhsFloat {
			return lhsFloat / rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.STAR:
		if isLhsFloat && isRhsFloat {
			return lhsFloat * rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.BANG_EQUAL:
		return lhs != rhs, nil
	case tokentype.EQUAL_EQUAL:
		return lhs == rhs, nil
	case tokentype.GREATER:
		if isLhsFloat && isRhsFloat {
			return lhsFloat > rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.GREATER_EQUAL:
		if isLhsFloat && isRhsFloat {
			return lhsFloat >= rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.LESS:
		if isLhsFloat && isRhsFloat {
			return lhsFloat < rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.LESS_EQUAL:
		if isLhsFloat && isRhsFloat {
			return lhsFloat <= rhsFloat, nil
		} else {
			return nil, loxerr.Runtime(e.Operator, invalidOperatorMsg)
		}
	case tokentype.AND:
		return conversion.IsTruthy(lhs) && conversion.IsTruthy(rhs), nil
	case tokentype.OR:
		return conversion.IsTruthy(lhs) || conversion.IsTruthy(rhs), nil
	default:
		return nil, loxerr.Internal(fmt.Sprintf("Unknown binary operator '%s' reached interpreter!", e.Operator.Lexeme))
	}
}

func (a *AstInterpreter) VisitUnaryExpr(e *ast.UnaryExpr) (interface{}, error) {
	rhsResult, err := e.Right.Accept(a)
	if err != nil {
		return nil, err
	}

	switch e.Operator.Type {
	case tokentype.BANG:
		return !conversion.IsTruthy(rhsResult), nil
	case tokentype.MINUS:
		if floatValue, ok := conversion.ToFloat(rhsResult); ok {
			return -floatValue, nil
		}
		return nil, loxerr.Runtime(e.Operator, fmt.Sprintf("Unable to apply operator '%s' to value: %v", e.Operator.Lexeme, rhsResult))
	default:
		return nil, loxerr.Internal(fmt.Sprintf("Unknown unary operator '%s' reached interpreter!", e.Operator.Lexeme))
	}
}

func (a *AstInterpreter) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	return e.Expression.Accept(a)
}

func (a *AstInterpreter) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	return e.Value, nil
}

func (a *AstInterpreter) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	result, err := a.findVar(e.Name, e)
	return result, err
}

func (a *AstInterpreter) ExecuteBlock(stmts []ast.Stmt, env *environment.Environment) error {
	prevEnv := a.env
	defer func() {
		a.env = prevEnv
	}()

	a.env = env

	for _, stmt := range stmts {
		_, err := stmt.Accept(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AstInterpreter) Resolve(expr ast.Expr, distance int) {
	a.localResolution[expr] = distance
}

func (a *AstInterpreter) findVar(name *token.Token, expr ast.Expr) (interface{}, error) {
	if distance, ok := a.localResolution[expr]; ok {
		return a.env.GetAt(name.Lexeme, distance), nil
	}

	result, exists := a.globals.Get(name.Lexeme)
	if !exists {
		return nil, loxerr.Runtime(name, fmt.Sprintf("Variable '%s' not defined", name.Lexeme))
	}
	return result, nil
}
