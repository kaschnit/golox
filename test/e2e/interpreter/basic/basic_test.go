package basic_test

import (
	"os"
	"testing"

	"github.com/kaschnit/golox/test/e2e/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testutil.BuildTestBinary()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestOutput_HelloWorld(t *testing.T) {
	result, err := testutil.InterpretTestProgram("basic/HelloWorld.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", result)
}

func TestOutput_PowerOfTwo(t *testing.T) {
	result, err := testutil.InterpretTestProgram("basic/PowerOfTwo.lox")
	assert.Nil(t, err)
	assert.Equal(t, "32", result)
}

func TestOutput_RecursiveFactorial(t *testing.T) {
	result, err := testutil.InterpretTestProgram("basic/RecursiveFactorial.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Recursive Factorials: 0! = 1; 1! = 1; 2! = 2; 3! = 6; 4! = 24; 5! = 120", result)
}
