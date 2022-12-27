package interpreter

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/environment"
)

type Callable interface {
	Arity() int
	Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error)
}

type LoxFunction struct {
	declaration *ast.FunctionStmt
	closure     *environment.Environment
}

func NewLoxFunction(declaration *ast.FunctionStmt, closure *environment.Environment) *LoxFunction {
	return &LoxFunction{
		declaration: declaration,
		closure:     closure,
	}
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	env := f.closure.NewChild()
	for i, param := range f.declaration.Params {
		env.Set(param.Lexeme, args[i])
	}

	err := interpreter.ExecuteBlock(f.declaration.Body, env)

	// Return is propagated by child nodes up until this node
	// to end execution of the function.
	if returnWrapper, ok := err.(*Return); ok {
		return returnWrapper.value, nil
	}

	return nil, err
}

type NativeFunction struct {
	arity int
	code  func(interpreter *AstInterpreter, args []interface{}) (interface{}, error)
}

func NewNativeFunction(arity int, code func(interpreter *AstInterpreter, args []interface{}) (interface{}, error)) *NativeFunction {
	return &NativeFunction{arity, code}
}

func (f *NativeFunction) Arity() int {
	return f.arity
}

func (f *NativeFunction) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	return f.code(interpreter, args)
}
