package interpreter

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/ast/analyzer"
	"github.com/kaschnit/golox/pkg/ast/astutil"
)

type InterpreterWrapper struct {
	analyzer    *analyzer.AstAnalyzer
	interpreter *AstInterpreter
}

func NewInterpreterWrapper() *InterpreterWrapper {
	return &InterpreterWrapper{
		analyzer:    analyzer.NewAstAnalyzer(),
		interpreter: NewAstInterpreter(),
	}
}

func (w *InterpreterWrapper) visitors() []ast.AstVisitor {
	return []ast.AstVisitor{
		w.analyzer,
		w.interpreter,
	}
}

func (w *InterpreterWrapper) InterpretSourceFile(filepath string) error {
	return astutil.ParseSourceFileAndVisit(filepath, w.visitors()...)
}

func (w *InterpreterWrapper) InterpretLine(line string) error {
	return astutil.ParseLineAndVisit(line, w.visitors()...)
}
