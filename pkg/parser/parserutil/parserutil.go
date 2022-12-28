package parserutil

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/scanner/scannerutil"
)

func ParseSourceFile(filepath string) (*ast.Program, error) {
	tokens, err := scannerutil.ScanSourceFile(filepath)
	if err != nil {
		return nil, err
	}

	// Parse the input.
	parser := parser.NewParser(tokens)
	programAst, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	return programAst, nil
}
