package repl_common

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/parser"
	"github.com/kaschnit/golox/pkg/scanner"
)

func ParseLineAndVisit(visitor ast.AstVisitor, line string) {
	// Tokenize the input.
	scanner := scanner.NewScanner(line)
	tokens, errs := scanner.ScanAllTokens()
	if len(errs) > 0 {
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return
	}

	// Parse the input.
	parser := parser.NewParser(tokens)
	programAst, errs := parser.Parse()
	if len(errs) > 0 {
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return
	}

	// Print the AST.
	programAst.Accept(visitor)
}
