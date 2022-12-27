package interpreter

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/interpreter/environment"
)

type Callable interface {
	Arity() int
	Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error)
}

// Runtime representation of user-defined function
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
		return returnWrapper.Value, nil
	}

	return nil, err
}

func (f *LoxFunction) String() string {
	return fmt.Sprintf("<function %s [%p]>", f.declaration.Name.Lexeme, f)
}

// Runtime representation of user-defined class
type LoxClass struct {
	declaration *ast.ClassStmt
}

func NewLoxClass(declaration *ast.ClassStmt) *LoxClass {
	return &LoxClass{declaration: declaration}
}

func (c *LoxClass) Arity() int {
	return 0
}

func (c *LoxClass) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	return NewLoxClassInstance(c), nil
}

func (c *LoxClass) String() string {
	return fmt.Sprintf("<class %s [%p]>", c.declaration.Name.Lexeme, c)
}

// Runtime representation of interpreter-defined ("native") function.
type NativeFunction struct {
	name  string
	arity int
	code  func(interpreter *AstInterpreter, args []interface{}) (interface{}, error)
}

func NewNativeFunction(name string, arity int, code func(interpreter *AstInterpreter, args []interface{}) (interface{}, error)) *NativeFunction {
	return &NativeFunction{
		name:  name,
		arity: arity,
		code:  code,
	}
}

func (f *NativeFunction) Arity() int {
	return f.arity
}

func (f *NativeFunction) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	return f.code(interpreter, args)
}

func (f *NativeFunction) String() string {
	return fmt.Sprintf("<native function %s>", f.name)
}
