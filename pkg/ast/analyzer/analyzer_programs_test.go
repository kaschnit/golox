package analyzer

import (
	"testing"

	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/test/programs"
	"github.com/stretchr/testify/assert"
)

func assertProgramHasNoAnalyzerErrors(t *testing.T, programPath string) {
	filepath := programs.GetPath(programPath)
	programAst, err := parser.ParseSourceFile(filepath)
	assert.Nil(t, err)

	result, err := NewAstAnalyzer().VisitProgram(programAst)
	assert.Nil(t, result)
	assert.Nil(t, err)
}

func TestAnalyzer_ConstructProgram_Assignment(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/Assignment.lox")
}

func TestAnalyzer_ConstructProgram_ClassConstructorNoArgs(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/ClassConstructorNoArgs.lox")
}

func TestAnalyzer_ConstructProgram_ClassConstructorWithArgs(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/ClassConstructorWithArgs.lox")
}

func TestAnalyzer_ConstructProgram_ClassMethods(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/ClassMethods.lox")
}

func TestAnalyzer_ConstructProgram_ClassThisKeyword(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/ClassThisKeyword.lox")
}

func TestAnalyzer_ConstructProgram_ForLoop(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/ForLoop.lox")
}

func TestAnalyzer_ConstructProgram_FunctionCallWithArgs(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/FunctionCallWithArgs.lox")
}

func TestAnalyzer_ConstructProgram_GlobalClosure(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/GlobalClosure.lox")
}

func TestAnalyzer_ConstructProgram_IfElse(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/IfElse.lox")
}

func TestAnalyzer_ConstructProgram_IfElseIf(t *testing.T) {
	assertProgramHasNoAnalyzerErrors(t, "constructs/IfElseIf.lox")
}
