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
	paramActuals := make(map[string]interface{})
	for i, param := range f.declaration.Params {
		paramActuals[param.Lexeme] = args[i]
	}

	env := f.closure.WithValues(paramActuals)
	err := interpreter.ExecuteBlock(f.declaration.Body, env)

	// Return is propagated by child nodes up until this node
	// to end execution of the function.
	if returnWrapper, ok := err.(*Return); ok {
		return returnWrapper.Value, nil
	}

	return nil, err
}

func (f *LoxFunction) Bind(instance *LoxClassInstance) *LoxFunction {
	closure := f.closure.WithValue("this", instance)
	return NewLoxFunction(f.declaration, closure)
}

func (f *LoxFunction) String() string {
	return fmt.Sprintf("<function %s [%p]>", f.declaration.Name.Lexeme, f)
}

// Runtime representation of user-defined class
type LoxClass struct {
	declaration       *ast.ClassStmt
	closure           *environment.Environment
	methods           map[string]*LoxFunction
	metaclassInstance *LoxClassInstance
}

func NewLoxClass(declaration *ast.ClassStmt, closure *environment.Environment) *LoxClass {
	metaclass := &LoxClass{
		declaration:       declaration,
		closure:           closure,
		methods:           getFunctionsMap(declaration.StaticMethods, closure),
		metaclassInstance: nil,
	}
	metaclassInstance := NewLoxClassInstance(metaclass)

	return &LoxClass{
		declaration:       declaration,
		closure:           closure,
		methods:           getFunctionsMap(declaration.Methods, closure),
		metaclassInstance: metaclassInstance,
	}
}

func (c *LoxClass) Arity() int {
	if c.declaration.Constructor == nil {
		return 0
	}
	return len(c.declaration.Constructor.Params)
}

func (c *LoxClass) Call(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
	instance := NewLoxClassInstance(c)

	// Call the constructor if it's been defined
	if c.declaration.Constructor != nil {
		constructor := NewLoxFunction(c.declaration.Constructor, c.closure)
		constructor.Bind(instance).Call(interpreter, args)
	}

	return instance, nil
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
	return fmt.Sprintf("<native function %s [%p]>", f.name, f)
}

func getFunctionsMap(declarations []*ast.FunctionStmt, closure *environment.Environment) map[string]*LoxFunction {
	functions := make(map[string]*LoxFunction)
	for _, function := range declarations {
		functions[function.Name.Lexeme] = NewLoxFunction(function, closure)
	}
	return functions
}
