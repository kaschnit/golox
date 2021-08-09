package parser

import (
	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/scanner"
)

type Parser struct {
	scanner *scanner.Scanner
}

func NewParser(scanner *scanner.Scanner) *Parser {
	return &Parser{scanner}
}

func (p *Parser) Parse() *ast.Program {
	// TODO
	return &ast.Program{Statements: []ast.Stmt{}}
}
