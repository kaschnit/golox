package constructs_test

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

func TestOutput_Construct_Assignment(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/Assignment.lox")
	assert.Nil(t, err)
	assert.Equal(t, "-12 398", result)
}

func TestOutput_Construct_ClassArgsConstructor(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/ClassArgsConstructor.lox")
	assert.Nil(t, err)
	assert.Equal(t, "bye bob bar", result)

}

func TestOutput_Construct_ClassNoArgsConstructor(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/ClassNoArgsConstructor.lox")
	assert.Nil(t, err)
	assert.Equal(t, "a1a2a3", result)
}

func TestOutput_Construct_ForLoop(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/ForLoop.lox")
	assert.Nil(t, err)
	assert.Equal(t, "0 1 2 3 4 5 Text Text ", result)
}

func TestOutput_Construct_FunctionCallWithArgs(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/FunctionCallWithArgs.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Printing 0: Printing 1: a Printing 2: b c Printing 3: d e f", result)
}

func TestOutput_Construct_GlobalClosure(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/GlobalClosure.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1112 334", result)
}

func TestOutput_Construct_IfElse(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/IfElse.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 if 2 if 3 else 4 if 5 else 6 if 7 else 8 if ", result)
}

func TestOutput_Construct_IfElseIf(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/IfElseIf.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 if 2 else if 3 else ", result)
}

func TestOutput_Construct_LocalClosure(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/LocalClosure.lox")
	assert.Nil(t, err)
	assert.Equal(t, "HelloHello15", result)
}

func TestOutput_Construct_LogicalAnd(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/LogicalAnd.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false false false true false", result)
}

func TestOutput_Construct_LogicalOr(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/LogicalOr.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false true true true true", result)
}

func TestOutput_Construct_NumericArithmeticOperations(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/NumericArithmeticOperations.lox")
	assert.Nil(t, err)
	assert.Equal(t, "3 -13 60 7.5 2", result)
}

func TestOutput_Construct_NumericComparisonOperations(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/NumericComparisonOperations.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false true false false true true true false false true false true true false false true", result)
}

func TestOutput_Construct_ReturnAtEndOfFunction(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/ReturnAtEndOfFunction.lox")
	assert.Nil(t, err)
	assert.Equal(t, "This should be printed! Yay!", result)
}

func TestOutput_Construct_ReturnEarlyFromFunction(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/ReturnEarlyFromFunction.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Yay!", result)
}

func TestOutput_Construct_Scoping(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/Scoping.lox")
	assert.Nil(t, err)
	assert.Equal(t, "10 9 5 6 9", result)
}

func TestOutput_Construct_WhileLoop(t *testing.T) {
	result, err := testutil.InterpretTestProgram("constructs/WhileLoop.lox")
	assert.Nil(t, err)
	assert.Equal(t, "4 3 2 1 0 ", result)
}
