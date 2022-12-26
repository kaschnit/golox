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
	return &LoxFunction{declaration, closure}
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	env := f.closure.Fork()
	for i := range f.declaration.Params {
		env.Set(f.declaration.Params[i].Lexeme, args[i])
	}
	return nil, interpreter.ExecuteBlock(f.declaration.Body, env)
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
