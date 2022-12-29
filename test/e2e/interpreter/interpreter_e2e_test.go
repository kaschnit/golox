package basic_test

import (
	"os"
	"testing"

	"github.com/kaschnit/golox/test/e2e/e2e_testutil"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	e2e_testutil.BuildTestBinary()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestOutput_HelloWorld(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("basic/HelloWorld.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", result)
}

func TestOutput_PowerOfTwo(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("basic/PowerOfTwo.lox")
	assert.Nil(t, err)
	assert.Equal(t, "32", result)
}

func TestOutput_RecursiveFactorial(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("basic/RecursiveFactorial.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Recursive Factorials: 0! = 1; 1! = 1; 2! = 2; 3! = 6; 4! = 24; 5! = 120", result)
}

func TestOutput_Construct_IfElseIf(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("constructs/IfElseIf.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 if 2 else if 3 else ", result)
}

func TestOutput_Construct_Scoping(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("constructs/Scoping.lox")
	assert.Nil(t, err)
	assert.Equal(t, "10 9 5 6 9", result)
}

func TestOutput_Construct_WhileLoop(t *testing.T) {
	result, err := e2e_testutil.InterpretTestProgram("constructs/WhileLoop.lox")
	assert.Nil(t, err)
	assert.Equal(t, "4 3 2 1 0 ", result)
}
