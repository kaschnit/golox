package parser

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/scanner"
)

func ParseSourceFile(filepath string) (*ast.Program, error) {
	tokens, err := scanner.ScanSourceFile(filepath)
	if err != nil {
		return nil, err
	}

	// Parse the input.
	parser := NewParser(tokens)
	programAst, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	return programAst, nil
}
