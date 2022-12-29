package interpreter

import (
	"testing"

	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/test/programs"
	"github.com/stretchr/testify/assert"
)

func TestInterpreter_InvalidProgram_GetClassInstanceUndefinedProperty(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/GetClassInstanceUndefinedProperty.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "hello", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "Property")
	assert.ErrorContains(t, runtimeErr, "not defined")
}
func TestInterpreter_InvalidProgram_PropertyAccessOnNonClass(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/PropertyAccessOnNonClass.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "someProperty", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "class instance")
}

func TestInterpreter_InvalidProgram_PropertySetOnNonClass(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/PropertySetOnNonClass.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "someProperty", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "class instance")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedAssignment(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/VariableNotDefinedAssignment.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "y", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedInitialization(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/VariableNotDefinedInitialization.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "y", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}

func TestInterpreter_InvalidProgram_VariableNotDefinedPrint(t *testing.T) {
	filepath := programs.GetPath("invalid/interpreter/VariableNotDefinedPrint.lox")
	err := InterpretSourceFile(filepath)
	assert.Error(t, err)
	assert.IsType(t, &loxerr.LoxRuntimeError{}, err)

	runtimeErr := err.(*loxerr.LoxRuntimeError)
	assert.Equal(t, "x", runtimeErr.Token.Lexeme)
	assert.ErrorContains(t, runtimeErr, "not defined")
}
