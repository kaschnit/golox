package parser

import (
	"github.com/kaschnit/golox/pkg/ast"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type Parser struct {
	tokens  []*token.Token
	start   int
	current int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		start:   0,
		current: 0,
	}
}

func (p *Parser) Parse() (*ast.Program, []error) {
	if numTokens := len(p.tokens); numTokens == 0 {
		return nil, []error{loxerr.NewLoxErrorAtLine(0, "Expected EOF.")}
	} else if p.tokens[numTokens-1].Type != tokentype.EOF {
		return nil, []error{loxerr.NewLoxErrorAtToken(p.tokens[numTokens-1], "Expected EOF.")}
	}
	return p.parseProgram()
}

func (p *Parser) parseProgram() (*ast.Program, []error) {
	errors := make([]error, 0)
	statements := make([]ast.Stmt, 0)
	for !p.isAtEnd() {
		stmt, err := p.parseStatement()
		if err != nil {
			errors = append(errors, err)
			p.synchronize()
		} else {
			statements = append(statements, stmt)
		}
	}
	return &ast.Program{Statements: statements}, errors
}

func (p *Parser) parseStatement() (ast.Stmt, error) {
	if p.peekMatches(1, tokentype.PRINT) {
		p.advance()
		return p.parsePrintStatement()
	}
	return p.parseExpressionStatement()
}

func (p *Parser) parsePrintStatement() (*ast.PrintStmt, error) {
	printExpr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	err = p.consume(tokentype.SEMICOLON, "Expected ';' after expression.")
	if err != nil {
		return nil, err
	}
	return &ast.PrintStmt{Expr: printExpr}, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExprStmt, error) {
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	err = p.consume(tokentype.SEMICOLON, "Expected ';' after expression.")
	if err != nil {
		return nil, err
	}
	return &ast.ExprStmt{Expr: expr}, nil
}

func (p *Parser) parseExpression() (ast.Expr, error) {
	return p.parseEquality()
}

func (p *Parser) parseEquality() (ast.Expr, error) {
	expr, err := p.parseComparison()
	if err != nil {
		return nil, err
	}

	for p.peekMatches(1, tokentype.BANG_EQUAL, tokentype.EQUAL_EQUAL) {
		left := expr
		operator := p.advance()
		right, err := p.parseComparison()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) parseComparison() (ast.Expr, error) {
	expr, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	matchTokens := []tokentype.TokenType{
		tokentype.GREATER, tokentype.GREATER_EQUAL,
		tokentype.LESS, tokentype.LESS_EQUAL,
	}
	for p.peekMatches(1, matchTokens...) {
		left := expr
		operator := p.advance()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) parseTerm() (ast.Expr, error) {
	expr, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for p.peekMatches(1, tokentype.MINUS, tokentype.PLUS) {
		left := expr
		operator := p.advance()
		right, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) parseFactor() (ast.Expr, error) {
	expr, err := p.parseUnary()
	if err != nil {
		return nil, err
	}
	for p.peekMatches(1, tokentype.SLASH, tokentype.STAR) {
		left := expr
		operator := p.advance()
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) parseUnary() (ast.Expr, error) {
	if p.peekMatches(1, tokentype.BANG, tokentype.MINUS) {
		operator := p.advance()
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{
			Operator: operator,
			Right:    right,
		}, nil
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (ast.Expr, error) {
	matchTokens := []tokentype.TokenType{
		tokentype.NUMBER, tokentype.STRING,
		tokentype.TRUE, tokentype.FALSE,
		tokentype.NIL,
	}
	if p.peekMatches(1, matchTokens...) {
		matched := p.advance()
		switch matched.Type {
		case tokentype.TRUE:
			return &ast.LiteralExpr{Value: true}, nil
		case tokentype.FALSE:
			return &ast.LiteralExpr{Value: false}, nil
		default:
			return &ast.LiteralExpr{Value: matched.Literal}, nil
		}
	} else if p.peekMatches(1, tokentype.LEFT_PAREN) {
		p.advance()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		err = p.consume(tokentype.RIGHT_PAREN, "Expected ')' after expression.")
		if err != nil {
			return nil, err
		}
		return &ast.GroupingExpr{Expression: expr}, nil
	} else {
		return nil, loxerr.NewLoxErrorAtToken(p.peek(1), "Expected expression.")
	}
}

func (p *Parser) consume(typeToMatch tokentype.TokenType, errorMessage string) error {
	nextToken := p.peek(1)
	if nextToken.Type == typeToMatch {
		p.current++
		return nil
	}
	return loxerr.NewLoxErrorAtToken(nextToken, errorMessage)
}

func (p *Parser) advance() *token.Token {
	result := p.peek(1)
	if !p.isAtEnd() {
		p.current++
	}
	return result
}

func (p *Parser) peekMatches(lookahead int, tokenTypes ...tokentype.TokenType) bool {
	for i := 0; i < len(tokenTypes); i++ {
		if p.peek(lookahead).Type == tokenTypes[i] {
			return true
		}
	}
	return false
}

func (p *Parser) peek(lookahead int) *token.Token {
	return p.tokens[p.current+lookahead-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek(1).Type == tokentype.EOF
}

func (p *Parser) synchronize() {
	if p.peekMatches(1, tokentype.SEMICOLON) {
		return
	}

	for !p.isAtEnd() {
		switch p.peek(1).Type {
		case tokentype.CLASS:
			fallthrough
		case tokentype.FUN:
			fallthrough
		case tokentype.VAR:
			fallthrough
		case tokentype.FOR:
			fallthrough
		case tokentype.IF:
			fallthrough
		case tokentype.WHILE:
			fallthrough
		case tokentype.PRINT:
			fallthrough
		case tokentype.RETURN:
			return
		}
		p.current++
	}
}
