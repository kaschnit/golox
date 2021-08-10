package parser

import (
	"github.com/kaschnit/golox/pkg/ast"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type Parser struct {
	tokens  []*token.Token
	errors  []error
	start   int
	current int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		start:   0,
		current: 0,
		errors:  make([]error, 0),
	}
}

func (p *Parser) Parse() *ast.Program {
	// TODO implement
	return nil
}

func (p *Parser) parseProgram() ast.Expr {
	// TODO implement
	return nil
}

func (p *Parser) parseStatement() ast.Expr {
	// TODO implement
	return nil
}

func (p *Parser) parseExpression() ast.Expr {
	return p.parseEquality()
}

func (p *Parser) parseEquality() ast.Expr {
	expr := p.parseComparison()
	for p.peekMatches(1, tokentype.BANG_EQUAL, tokentype.EQUAL_EQUAL) {
		left := expr
		operator := p.advance()
		right := p.parseComparison()
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) parseComparison() ast.Expr {
	matchTokens := []tokentype.TokenType{
		tokentype.GREATER, tokentype.GREATER_EQUAL,
		tokentype.LESS, tokentype.LESS_EQUAL,
	}

	expr := p.parseTerm()
	for p.peekMatches(1, matchTokens...) {
		left := expr
		operator := p.advance()
		right := p.parseTerm()
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) parseTerm() ast.Expr {
	expr := p.parseFactor()
	for p.peekMatches(1, tokentype.MINUS, tokentype.PLUS) {
		left := expr
		operator := p.advance()
		right := p.parseFactor()
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) parseFactor() ast.Expr {
	expr := p.parseUnary()
	for p.peekMatches(1, tokentype.SLASH, tokentype.STAR) {
		left := expr
		operator := p.advance()
		right := p.parseUnary()
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) parseUnary() ast.Expr {
	if p.peekMatches(1, tokentype.BANG, tokentype.MINUS) {
		operator := p.advance()
		right := p.parseUnary()
		return &ast.UnaryExpr{
			Operator: operator,
			Right:    right,
		}
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() ast.Expr {
	matchTokens := []tokentype.TokenType{
		tokentype.NUMBER, tokentype.STRING,
		tokentype.TRUE, tokentype.FALSE,
		tokentype.NIL,
	}
	if p.peekMatches(1, matchTokens...) {
		matched := p.advance()
		switch matched.Type {
		case tokentype.TRUE:
			return &ast.LiteralExpr{Value: true}
		case tokentype.FALSE:
			return &ast.LiteralExpr{Value: false}
		default:
			return &ast.LiteralExpr{Value: matched.Literal}
		}
	} else if p.peekMatches(1, tokentype.LEFT_PAREN) {
		p.advance()
		result := p.parseExpression()
		err := p.consume(tokentype.RIGHT_PAREN, "Expected ')' after expression.")
		if err != nil {
			p.errors = append(p.errors, err)
		}
		return result
	} else {
		// This means there is a bug in the parser, not a bad input.
		err := loxerr.NewLoxInternalError("No tokens matched by parser")
		panic(err)
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
	p.current++
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
