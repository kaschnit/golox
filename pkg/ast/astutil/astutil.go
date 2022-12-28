package astutil

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/parser/parserutil"
	"github.com/kaschnit/golox/pkg/scanner"
)

// Parse the source code in the file located at filepath and apply the visitor to the root
// of the AST that is produced, visiting each node in the AST.
func ParseSourceFileAndVisit(filepath string, visitors ...ast.AstVisitor) error {
	programAst, err := parserutil.ParseSourceFile(filepath)
	if err != nil {
		return err
	}

	for _, visitor := range visitors {
		_, err = visitor.VisitProgram(programAst)
		if err != nil {
			return err
		}
	}

	return nil
}

// Parse the line of source code and apply the visitor to the root
// of the AST that is produced, visiting each node in the AST.
func ParseLineAndVisit(line string, visitors ...ast.AstVisitor) error {
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
	for _, visitor := range visitors {
		_, err = programAst.Accept(visitor)
		if err != nil {
			return err
		}
	}

	return nil
}
