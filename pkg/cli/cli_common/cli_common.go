package cli_common

import (
	"io"
	"os"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/scanner"
)

func ParseSourceFileAndVisit(visitor ast.AstVisitor, filepath string) error {
	f, err := os.Open(filepath)
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return err
	}

	sourceCode, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	scanner := scanner.NewScanner(string(sourceCode))
	tokens, err := scanner.ScanAllTokens()
	if err != nil {
		return err
	}

	// Parse the input.
	parser := parser.NewParser(tokens)
	programAst, err := parser.Parse()
	if err != nil {
		return err
	}

	_, err = visitor.VisitProgram(programAst)
	if err != nil {
		return err
	}

	return nil
}

func ParseLineAndVisit(visitor ast.AstVisitor, line string) error {
	// Tokenize the input.
	scanner := scanner.NewScanner(line)
	tokens, err := scanner.ScanAllTokens()
	if err != nil {
		return err
	}

	// Parse the input.
	parser := parser.NewParser(tokens)
	programAst, err := parser.Parse()
	if err != nil {
		return err
	}

	// Visit the AST.
	_, err = programAst.Accept(visitor)
	if err != nil {
		return err
	}

	return nil
}
