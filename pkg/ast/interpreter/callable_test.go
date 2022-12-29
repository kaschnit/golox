package interpreter

import (
	"fmt"
	"testing"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/interpreter/environment"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestLoxClass_ToString(t *testing.T) {
	clsDecl := &ast.ClassStmt{Name: &token.Token{Lexeme: ""}}
	env := environment.NewEnvironment(make(map[string]interface{}))
	cls := NewLoxClass(clsDecl, env)

	var result string

	clsDecl.Name.Lexeme = "MyClass"
	result = cls.String()
	assert.Contains(t, result, "<class MyClass")
	assert.Contains(t, result, ">")

	clsDecl.Name.Lexeme = "SomeOtherName"
	result = cls.String()
	assert.Contains(t, result, "<class SomeOtherName")
	assert.Contains(t, result, ">")
}

func TestLoxClass_Arity(t *testing.T) {
	clsDecl := &ast.ClassStmt{Name: &token.Token{Lexeme: "MyClass"}}
	env := environment.NewEnvironment(make(map[string]interface{}))
	cls := NewLoxClass(clsDecl, env)
	assert.Equal(t, 0, cls.Arity())

	clsDecl.Constructor = &ast.FunctionStmt{
		Params: []*token.Token{{}, {}},
	}
	assert.Equal(t, 2, cls.Arity())

	clsDecl.Constructor = nil
	assert.Equal(t, 0, cls.Arity())

	clsDecl.Constructor = &ast.FunctionStmt{
		Params: []*token.Token{},
	}
	assert.Equal(t, 0, cls.Arity())
}

func TestLoxFunction_Arity(t *testing.T) {
	funcDecl := &ast.FunctionStmt{Params: []*token.Token{}}
	env := environment.NewEnvironment(make(map[string]interface{}))
	loxFunc := NewLoxFunction(funcDecl, env)
	assert.Equal(t, 0, loxFunc.Arity())

	funcDecl.Params = []*token.Token{{}, {}, {}}
	assert.Equal(t, 3, loxFunc.Arity())
}

func TestLoxFunction_String(t *testing.T) {
	funcDecl := &ast.FunctionStmt{Name: &token.Token{Lexeme: "myFunc1"}}
	env := environment.NewEnvironment(make(map[string]interface{}))
	loxFunc := NewLoxFunction(funcDecl, env)

	result := loxFunc.String()
	assert.Contains(t, result, "<function myFunc1 [")
	assert.Contains(t, result, "]>")
}

func TestNativeFunction_Arity(t *testing.T) {
	nativeFunc := NewNativeFunction(
		"myAwesomeFunction",
		0,
		func(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
			return 1, nil
		},
	)
	assert.Equal(t, 0, nativeFunc.Arity())

	nativeFunc = NewNativeFunction(
		"myOtherAwesomeFunction",
		12,
		func(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
			return 100, nil
		},
	)
	assert.Equal(t, 12, nativeFunc.Arity())
}

func TestNativeFunction_Call(t *testing.T) {
	counter := 0
	nativeFunc := NewNativeFunction(
		"myAwesomeFunction",
		0,
		func(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
			counter += 1
			return fmt.Sprintf("Hello %d (%v)", counter, args[0]), nil
		},
	)

	args := []interface{}{"abc"}
	result, err := nativeFunc.Call(nil, args)
	assert.Nil(t, err)
	assert.Equal(t, "Hello 1 (abc)", result)
	assert.Equal(t, 1, counter)

	args = []interface{}{"abc"}
	result, err = nativeFunc.Call(nil, args)
	assert.Nil(t, err)
	assert.Equal(t, "Hello 2 (abc)", result)
	assert.Equal(t, 2, counter)

	args = []interface{}{"defghijk"}
	result, err = nativeFunc.Call(nil, args)
	assert.Nil(t, err)
	assert.Equal(t, "Hello 3 (defghijk)", result)
	assert.Equal(t, 3, counter)
}

func TestNativeFunction_ToString(t *testing.T) {
	nativeFunc := NewNativeFunction(
		"myAwesomeFunction",
		0,
		func(interpreter *AstInterpreter, args []interface{}) (interface{}, error) {
			return 1, nil
		},
	)

	result := nativeFunc.String()
	assert.Contains(t, result, "<native function myAwesomeFunction")
	assert.Contains(t, result, ">")
}
