package interpreter

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/conversion"
	"github.com/kaschnit/golox/pkg/environment"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type AstInterpreter struct {
	env *environment.Environment
}

func NewAstInterpreter() *AstInterpreter {
	return &AstInterpreter{
		env: environment.NewEnvironment(nil),
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
	currEnv := a.env
	a.setEnv(a.env.Fork())
	defer a.setEnv(currEnv)

	for i := 0; i < len(s.Statements); i++ {
		_, err := s.Statements[i].Accept(a)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (a *AstInterpreter) VisitVarStmt(s *ast.VarStmt) (interface{}, error) {
	varName := s.Left.Lexeme
	_, exists := a.env.Get(varName)
	if exists {
		return nil, loxerr.NewLoxErrorAtToken(s.Left, fmt.Sprintf("Variable '%s' already defined", varName))
	}
	value, err := s.Right.Accept(a)
	if err != nil {
		return nil, err
	}

	a.env.Set(varName, value)
	return nil, nil
}

func (a *AstInterpreter) VisitAssignExpr(e *ast.AssignExpr) (interface{}, error) {
	varName := e.Left.Lexeme
	_, exists := a.env.GetTraverse(varName)
	if !exists {
		return nil, loxerr.NewLoxErrorAtToken(e.Left, fmt.Sprintf("Variable '%s' not defined", varName))
	}

	value, err := e.Right.Accept(a)
	if err != nil {
		return nil, err
	}

	a.env.ReplaceTraverse(varName, value)
	return value, nil
}

func (p *AstInterpreter) VisitCallExpr(e *ast.CallExpr) (interface{}, error) {
	return nil, nil
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
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.PLUS:
		if isLhsFloat && isRhsFloat {
			return lhsFloat + rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.SLASH:
		if isLhsFloat && isRhsFloat {
			return lhsFloat / rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.STAR:
		if isLhsFloat && isRhsFloat {
			return lhsFloat * rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.BANG_EQUAL:
		return lhs != rhs, nil
	case tokentype.EQUAL_EQUAL:
		return lhs == rhs, nil
	case tokentype.GREATER:
		if isLhsFloat && isRhsFloat {
			return lhsFloat > rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.GREATER_EQUAL:
		if isLhsFloat && isRhsFloat {
			return lhsFloat >= rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.LESS:
		if isLhsFloat && isRhsFloat {
			return lhsFloat < rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.LESS_EQUAL:
		if isLhsFloat && isRhsFloat {
			return lhsFloat <= rhsFloat, nil
		} else {
			return nil, loxerr.NewLoxErrorAtToken(e.Operator, invalidOperatorMsg)
		}
	case tokentype.AND:
		return conversion.IsTruthy(lhs) && conversion.IsTruthy(rhs), nil
	case tokentype.OR:
		return conversion.IsTruthy(lhs) || conversion.IsTruthy(rhs), nil
	default:
		return nil, loxerr.NewLoxInternalError(fmt.Sprintf("Unknown binary operator '%s' reached interpreter!", e.Operator.Lexeme))
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
		return nil, loxerr.NewLoxErrorAtToken(e.Operator, fmt.Sprintf("Unable to apply operator '%s' to value: %v", e.Operator.Lexeme, rhsResult))
	default:
		return nil, loxerr.NewLoxInternalError(fmt.Sprintf("Unknown unary operator '%s' reached interpreter!", e.Operator.Lexeme))
	}
}

func (a *AstInterpreter) VisitGroupingExpr(e *ast.GroupingExpr) (interface{}, error) {
	return e.Expression.Accept(a)
}

func (a *AstInterpreter) VisitLiteralExpr(e *ast.LiteralExpr) (interface{}, error) {
	return e.Value, nil
}

func (a *AstInterpreter) VisitVarExpr(e *ast.VarExpr) (interface{}, error) {
	result, exists := a.env.GetTraverse(e.Name.Lexeme)
	if !exists {
		return nil, loxerr.NewLoxErrorAtToken(e.Name, fmt.Sprintf("Variable '%s' not defined", e.Name.Lexeme))
	}
	return result, nil
}

func (a *AstInterpreter) setEnv(env *environment.Environment) {
	a.env = env
}
