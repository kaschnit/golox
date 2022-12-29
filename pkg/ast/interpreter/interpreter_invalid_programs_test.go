package interpreter

import (
	"testing"

	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestInterpreter_InvalidProgram_GetClassInstanceUndefinedProperty(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/GetClassInstanceUndefinedProperty.lox")
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "hello", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "Property")
	assert.ErrorContains(t, runtimeErr, "not defined")
}
func TestInterpreter_InvalidProgram_PropertyAccessOnNonClass(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/PropertyAccessOnNonClass.lox")
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "someProperty", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "instance")
}

func TestInterpreter_InvalidProgram_PropertySetOnNonClass(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/PropertySetOnNonClass.lox")

	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "someProperty", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "instance")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedAssignment(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/VariableNotDefinedAssignment.lox")
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "y", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedInitialization(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/VariableNotDefinedInitialization.lox")
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "y", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedPrint(t *testing.T) {
	err := interpretSourceFile("invalid/interpreter/VariableNotDefinedPrint.lox")
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "x", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}
