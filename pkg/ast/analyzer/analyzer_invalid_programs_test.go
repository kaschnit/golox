package analyzer

import (
	"testing"

	"github.com/hashicorp/go-multierror"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/token/tokentype"
	"github.com/kaschnit/golox/test/programs"
	"github.com/stretchr/testify/assert"
)

func analyzeProgram(t *testing.T, subPath string) (interface{}, error) {
	filepath := programs.GetPath(subPath)
	programAst, err := parser.ParseSourceFile(filepath)
	assert.Nil(t, err)

	return NewAstAnalyzer().VisitProgram(programAst)
}

func TestAnalyzer_InvalidProgram_ReadLocalVariableInInitializer(t *testing.T) {
	result, err := analyzeProgram(t, "invalid/analyzer/ReadLocalVariableInInitializer.lox")
	assert.Nil(t, result)
	assert.IsType(t, &multierror.Error{}, err)

	errs := err.(*multierror.Error).Errors
	assert.Len(t, errs, 4)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[0])
	assert.Equal(t, tokentype.IDENTIFIER, errs[0].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.Equal(t, "y", errs[0].(*loxerr.LoxErrorAtToken).Token.Lexeme)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[1])
	assert.Equal(t, tokentype.IDENTIFIER, errs[1].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.Equal(t, "z", errs[1].(*loxerr.LoxErrorAtToken).Token.Lexeme)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[2])
	assert.Equal(t, tokentype.IDENTIFIER, errs[2].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.Equal(t, "a", errs[2].(*loxerr.LoxErrorAtToken).Token.Lexeme)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[3])
	assert.Equal(t, tokentype.IDENTIFIER, errs[3].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.Equal(t, "b", errs[3].(*loxerr.LoxErrorAtToken).Token.Lexeme)
}

func TestAnalyzer_InvalidProgram_ReturnValueInClassConstructor(t *testing.T) {
	result, err := analyzeProgram(t, "invalid/analyzer/ReturnValueInClassConstructor.lox")
	assert.Nil(t, result)
	assert.IsType(t, &multierror.Error{}, err)

	errs := err.(*multierror.Error).Errors
	assert.Len(t, errs, 4)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[0])
	assert.Equal(t, tokentype.RETURN, errs[0].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[0], "return")

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[1])
	assert.Equal(t, tokentype.RETURN, errs[1].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[1], "return")

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[2])
	assert.Equal(t, tokentype.RETURN, errs[1].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[2], "return")

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[3])
	assert.Equal(t, tokentype.RETURN, errs[1].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[3], "return")
}

func TestAnalyzer_InvalidProgram_ThisOutsideOfClass(t *testing.T) {
	result, err := analyzeProgram(t, "invalid/analyzer/ThisOutsideOfClass.lox")
	assert.Nil(t, result)
	assert.IsType(t, &multierror.Error{}, err)

	errs := err.(*multierror.Error).Errors
	assert.Len(t, errs, 2)

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[0])
	assert.Equal(t, tokentype.THIS, errs[0].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[0], "this")

	assert.IsType(t, &loxerr.LoxErrorAtToken{}, errs[1])
	assert.Equal(t, tokentype.THIS, errs[1].(*loxerr.LoxErrorAtToken).Token.Type)
	assert.ErrorContains(t, errs[1], "this")
}
