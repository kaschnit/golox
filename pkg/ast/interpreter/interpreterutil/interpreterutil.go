package interpreterutil

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/analyzer"
	"github.com/kaschnit/golox/pkg/ast/astutil"
	"github.com/kaschnit/golox/pkg/ast/interpreter"
)

func getVisitors() []ast.AstVisitor {
	return []ast.AstVisitor{
		analyzer.NewAstAnalyzer(),
		interpreter.NewAstInterpreter(),
	}
}

func InterpretSourceFile(filepath string) error {
	return astutil.ParseSourceFileAndVisit(filepath, getVisitors()...)
}

func InterpretLine(line string) error {
	return astutil.ParseLineAndVisit(line, getVisitors()...)
}
