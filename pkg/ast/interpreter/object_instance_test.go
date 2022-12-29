package interpreter

import (
	"testing"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/interpreter/environment"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestLoxClassInstance_ToString(t *testing.T) {
	clsDecl := &ast.ClassStmt{Name: &token.Token{Lexeme: ""}}
	env := environment.NewEnvironment(make(map[string]interface{}))
	cls := NewLoxClass(clsDecl, env)
	instance := NewLoxClassInstance(cls)

	var result string

	clsDecl.Name.Lexeme = "MyClass"
	result = instance.String()
	assert.Contains(t, result, "<instance of <class MyClass")
	assert.Contains(t, result, ">")

	clsDecl.Name.Lexeme = "SomeOtherName"
	result = instance.String()
	assert.Contains(t, result, "<instance of <class SomeOtherName")
	assert.Contains(t, result, ">")
}
