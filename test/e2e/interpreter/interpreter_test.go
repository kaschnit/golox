package interpreter_test

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
	result, err := testutil.InterpretTestProgram("HelloWorld.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", result)
}

func TestOutput_PowerOfTwo(t *testing.T) {
	result, err := testutil.InterpretTestProgram("PowerOfTwo.lox")
	assert.Nil(t, err)
	assert.Equal(t, "32", result)
}
