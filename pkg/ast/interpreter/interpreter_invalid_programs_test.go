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
