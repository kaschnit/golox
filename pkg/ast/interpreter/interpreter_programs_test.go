package interpreter

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/kaschnit/golox/test/programs"
	"github.com/kaschnit/golox/test/testutil"
	"github.com/stretchr/testify/assert"
)

func interpretSourceFile(subPath string) error {
	filepath := programs.GetPath(subPath)
	return NewInterpreterWrapper().InterpretSourceFile(filepath)
}

func getInterpreterOutput(programName string) (string, error) {
	return testutil.CaptureOutput(func() error {
		return interpretSourceFile(programName)
	})
}

func TestOutput_Construct_Assignment(t *testing.T) {
	result, err := getInterpreterOutput("constructs/Assignment.lox")
	assert.Nil(t, err)
	assert.Equal(t, "-12 398", result)
}

func TestOutput_Construct_ClassConstructorEarlyReturn(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassConstructorEarlyReturn.lox")
	assert.Nil(t, err)
	assert.Equal(t, "123123123", result)
}

func TestOutput_Construct_ClassConstructorWithArgs(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassConstructorWithArgs.lox")
	assert.Nil(t, err)
	assert.Equal(t, "bye bob bar", result)
}

func TestOutput_Construct_ClassConstructorNoArgs(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassConstructorNoArgs.lox")
	assert.Nil(t, err)
	assert.Equal(t, "a1a2a3", result)
}

func TestOutput_Construct_ClassMethods(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassMethods.lox")
	assert.Nil(t, err)
	assert.Equal(t, "AH! 100", result)
}
func TestOutput_Construct_ClassStaticMethods(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassStaticMethods.lox")
	assert.Nil(t, err)
	assert.Equal(t, "getting instance; instance value 99; static print; param static print 49", result)
}

func TestOutput_Construct_ClassThisKeyword(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ClassThisKeyword.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 10 11   10 99 10   99 99 12   99 99 13   13 1 13", result)
}

func TestOutput_Construct_ForLoop(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ForLoop.lox")
	assert.Nil(t, err)
	assert.Equal(t, "0 1 2 3 4 5 Text Text ", result)
}

func TestOutput_Construct_FunctionCallWithArgs(t *testing.T) {
	result, err := getInterpreterOutput("constructs/FunctionCallWithArgs.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Printing 0: Printing 1: a Printing 2: b c Printing 3: d e f", result)
}

func TestOutput_Construct_GlobalClosure(t *testing.T) {
	result, err := getInterpreterOutput("constructs/GlobalClosure.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1112 334", result)
}

func TestOutput_Construct_IfElse(t *testing.T) {
	result, err := getInterpreterOutput("constructs/IfElse.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 if 2 if 3 else 4 if 5 else 6 if 7 else 8 if ", result)
}

func TestOutput_Construct_IfElseIf(t *testing.T) {
	result, err := getInterpreterOutput("constructs/IfElseIf.lox")
	assert.Nil(t, err)
	assert.Equal(t, "1 if 2 else if 3 else ", result)
}

func TestOutput_Construct_LocalClosure(t *testing.T) {
	result, err := getInterpreterOutput("constructs/LocalClosure.lox")
	assert.Nil(t, err)
	assert.Equal(t, "HelloHello15", result)
}

func TestOutput_Construct_LogicalAnd(t *testing.T) {
	result, err := getInterpreterOutput("constructs/LogicalAnd.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false false false true false", result)
}

func TestOutput_Construct_LogicalOr(t *testing.T) {
	result, err := getInterpreterOutput("constructs/LogicalOr.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false true true true true", result)
}

func TestOutput_Construct_NativeFunction_Clock(t *testing.T) {
	result, err := getInterpreterOutput("constructs/NativeFunction_Clock.lox")
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(result, "<native function clock ["))
	assert.Contains(t, result, " => ")

	parts := strings.Split(result, " => ")
	assert.Len(t, parts, 2)
	assert.True(t, strings.HasPrefix(string(parts[0]), "<native function clock ["))
	assert.True(t, strings.HasSuffix(string(parts[0]), "]>"))

	unixTimestamp, err := strconv.ParseInt(parts[1], 10, 64)
	assert.Nil(t, err)

	oneHour, err := time.ParseDuration("1h")
	assert.Nil(t, err)

	laterTimestamp := time.Now().Add(oneHour)
	programTimestamp := time.Unix(unixTimestamp, 0)
	assert.Greater(t, laterTimestamp, programTimestamp)
}

func TestOutput_Construct_NumericArithmeticOperations(t *testing.T) {
	result, err := getInterpreterOutput("constructs/NumericArithmeticOperations.lox")
	assert.Nil(t, err)
	assert.Equal(t, "3 -13 60 7.5 2", result)
}

func TestOutput_Construct_NumericComparisonOperations(t *testing.T) {
	result, err := getInterpreterOutput("constructs/NumericComparisonOperations.lox")
	assert.Nil(t, err)
	assert.Equal(t, "false true false false true true true false false true false true true false false true", result)
}

func TestOutput_Construct_ReturnAtEndOfFunction(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ReturnAtEndOfFunction.lox")
	assert.Nil(t, err)
	assert.Equal(t, "This should be printed! Yay!", result)
}

func TestOutput_Construct_ReturnEarlyFromFunction(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ReturnEarlyFromFunction.lox")
	assert.Nil(t, err)
	assert.Equal(t, "Yay!", result)
}

func TestOutput_Construct_ReturnNoValue(t *testing.T) {
	result, err := getInterpreterOutput("constructs/ReturnNoValue.lox")
	assert.Nil(t, err)
	assert.Equal(t, "0123", result)
}

func TestOutput_Construct_Scoping(t *testing.T) {
	result, err := getInterpreterOutput("constructs/Scoping.lox")
	assert.Nil(t, err)
	assert.Equal(t, "10 9 5 6 9", result)
}

func TestOutput_Construct_WhileLoop(t *testing.T) {
	result, err := getInterpreterOutput("constructs/WhileLoop.lox")
	assert.Nil(t, err)
	assert.Equal(t, "4 3 2 1 0 ", result)
}
