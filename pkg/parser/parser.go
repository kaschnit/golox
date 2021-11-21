package parser

import (
	"github.com/kaschnit/golox/pkg/ast"
	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type Parser struct {
	// The tokenized source code to be parsed.
	tokens []*token.Token

	// The start of the sequence of tokens currently being parsed.
	start int

	// The current token in the sequence of tokens currently being parsed.
	current int
}

// Create a Parser instance.
func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		start:   0,
		current: 0,
	}
}

// Reset the parser to the initial state.
func (p *Parser) Reset() {
	p.start = 0
	p.current = 0
}

// Parse the entire tokenized source code, converting it to an AST
// with a root of type ast.Program.
func (p *Parser) Parse() (*ast.Program, error) {
	if numTokens := len(p.tokens); numTokens == 0 {
		return nil, loxerr.AtLine(0, "Expected EOF.")
	} else if p.tokens[numTokens-1].Type != tokentype.EOF {
		return nil, loxerr.AtToken(p.tokens[numTokens-1], "Expected EOF.")
	}
	return p.parseProgram()
}

// Parse a program, which is the root of the AST.
func (p *Parser) parseProgram() (*ast.Program, error) {
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

	program := &ast.Program{Statements: statements}
	if len(errors) > 0 {
		return program, loxerr.Multi(errors)
	} else {
		return program, nil
	}
}

// Parse a statement with the type of statement based on the next token.
func (p *Parser) parseStatement() (ast.Stmt, error) {
	nextToken := p.peek(1)

	switch nextToken.Type {
	case tokentype.PRINT:
		p.advance()
		return p.parsePrintStatement()
	case tokentype.RETURN:
		p.advance()
		return p.parseReturnStatement()
	case tokentype.IF:
		p.advance()
		return p.parseIfStatement()
	case tokentype.WHILE:
		p.advance()
		return p.parseWhileStatement()
	case tokentype.FOR:
		p.advance()
		return p.parseForStatement()
	case tokentype.FUN:
		p.advance()
		return p.parseFunctionStatement()
	case tokentype.VAR:
		p.advance()
		return p.parseVarStatement()
	case tokentype.LEFT_BRACE:
		p.advance()
		return p.parseBlockStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Parse a print statement.
func (p *Parser) parsePrintStatement() (*ast.PrintStmt, error) {
	printExpr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(tokentype.SEMICOLON, "Expected ';' after expression.")
	if err != nil {
		return nil, err
	}
	return &ast.PrintStmt{Expression: printExpr}, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStmt, error) {
	var expr ast.Expr
	var err error
	if !p.peekMatches(1, tokentype.SEMICOLON) {
		expr, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(tokentype.SEMICOLON, "Expected ';' after expression.")
	if err != nil {
		return nil, err
	}

	return &ast.ReturnStmt{
		Expression: expr,
	}, nil
}

// Parse an expression statement.
func (p *Parser) parseExpressionStatement() (*ast.ExprStmt, error) {

	var expr ast.Expr
	var err error
	if p.peekMatches(1, tokentype.SEMICOLON) {
		expr = &ast.LiteralExpr{Value: nil}
	} else {
		expr, err = p.parseExpression()
	}
	if err != nil {
		return nil, err
	}

	_, err = p.consume(tokentype.SEMICOLON, "Expected ';' after expression.")
	if err != nil {
		return nil, err
	}
	return &ast.ExprStmt{Expression: expr}, nil
}

// Parse an if statement.
func (p *Parser) parseIfStatement() (*ast.IfStmt, error) {
	var err error

	// Parse the parenthesized condition.
	_, err = p.consume(tokentype.LEFT_PAREN, "Expected '(' after 'if'.")
	if err != nil {
		return nil, err
	}
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokentype.RIGHT_PAREN, "Expected ')' after condition.")
	if err != nil {
		return nil, err
	}

	thenStatement, err := p.parseStatement()
	if err != nil {
		return nil, err
	}

	var elseStatement ast.Stmt
	if p.peekMatches(1, tokentype.ELSE) {
		p.advance()
		elseStatement, err = p.parseStatement()
		if err != nil {
			return nil, err
		}
	} else {
		elseStatement = nil
	}

	return &ast.IfStmt{
		Condition:     condition,
		ThenStatement: thenStatement,
		ElseStatement: elseStatement,
	}, nil
}

// Parse a while loop statement.
func (p *Parser) parseWhileStatement() (*ast.WhileStmt, error) {
	var err error

	_, err = p.consume(tokentype.LEFT_PAREN, "Expected '(' after 'while'.")
	if err != nil {
		return nil, err
	}
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokentype.RIGHT_PAREN, "Expected ')' after condition.")
	if err != nil {
		return nil, err
	}

	loopStatement, err := p.parseStatement()
	if err != nil {
		return nil, err
	}

	return &ast.WhileStmt{
		Condition:     condition,
		LoopStatement: loopStatement,
	}, nil
}

// Parse a for loop statement.
// Desugars the for loop to a while loop. The following two are equivalent:
// 	1. for (int i = 0; i < 5; i++) { doSomething() }
// 	2. { int i = 0; while (i < 5) { { doSomething(); } i++; } }
// The while loop is placed inside its own block. The initializer is placed at the beginning
// of this block. The loop body is placed inside a nested block, and the increment is placed
// after this nested block.
func (p *Parser) parseForStatement() (ast.Stmt, error) {
	var err error
	var nextToken *token.Token

	_, err = p.consume(tokentype.LEFT_PAREN, "Expected '(' after 'for'.")
	if err != nil {
		return nil, err
	}

	// Parse the initializer statement, if there is one.
	// It can be a var declaration or an expression statement.
	nextToken = p.peek(1)
	var initializer ast.Stmt
	if nextToken.Type == tokentype.SEMICOLON {
		p.advance()
	} else if nextToken.Type == tokentype.VAR {
		p.advance()
		initializer, err = p.parseVarStatement()
	} else {
		initializer, err = p.parseExpressionStatement()
	}
	if err != nil {
		return nil, err
	}

	// Parse the condition, if there is one.
	nextToken = p.peek(1)
	var condition ast.Expr
	if nextToken.Type == tokentype.SEMICOLON {
		condition = &ast.LiteralExpr{Value: true}
	} else {
		condition, err = p.parseExpression()
	}
	if err != nil {
		return nil, err
	}

	// Needs to be a ';' after the condition.
	_, err = p.consume(tokentype.SEMICOLON, "Expect ';' after loop condition.")
	if err != nil {
		return nil, err
	}

	// Parse the increment, if there is one.
	nextToken = p.peek(1)
	var increment ast.Expr
	if nextToken.Type != tokentype.RIGHT_PAREN {
		increment, err = p.parseExpression()
	}
	if err != nil {
		return nil, err
	}

	// Needs to be a ')' before the loop body.
	_, err = p.consume(tokentype.RIGHT_PAREN, "Expected ')' after loop increment.")
	if err != nil {
		return nil, err
	}

	// Parse the loop body.
	loopBody, err := p.parseStatement()
	if err != nil {
		return nil, err
	}

	// If there's an increment, place it at the end of the loop body.
	if increment != nil {
		incrementStmt := &ast.ExprStmt{Expression: increment}
		loopBody = &ast.BlockStmt{
			Statements: []ast.Stmt{loopBody, incrementStmt},
		}
	}

	// Construct the while loop from the parsed expressions and statement.
	var result ast.Stmt
	result = &ast.WhileStmt{
		Condition:     condition,
		LoopStatement: loopBody,
	}

	// If there's an initializer, wrap the while statement in a block
	// with the initializer then the while loop.
	if initializer != nil {
		result = &ast.BlockStmt{
			Statements: []ast.Stmt{initializer, result},
		}
	}

	return result, nil
}

// Parse a block statement.
func (p *Parser) parseBlockStatement() (*ast.BlockStmt, error) {
	statements := make([]ast.Stmt, 0)
	for !p.peekMatches(1, tokentype.RIGHT_BRACE) && !p.isAtEnd() {
		statement, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}
	_, err := p.consume(tokentype.RIGHT_BRACE, "Expected '}' after block.")
	if err != nil {
		return nil, err
	}
	return &ast.BlockStmt{Statements: statements}, nil
}

// Parse a func statement.
func (p *Parser) parseFunctionStatement() (*ast.FunctionStmt, error) {
	// Function declaration starts with the function's name.
	symbol, err := p.consume(tokentype.IDENTIFIER, "Expected identifier after 'fun'.")
	if err != nil {
		return nil, err
	}

	// Parse the args within the parentheses.
	_, err = p.consume(tokentype.LEFT_PAREN, "Expected '('.")
	if err != nil {
		return nil, err
	}

	args := make([]*token.Token, 0)
	if p.peekMatches(1, tokentype.RIGHT_PAREN) {
		// No args, move on.
		p.advance()
	} else {
		// Parse the args.
		for {
			arg, err := p.consume(tokentype.IDENTIFIER, "Expected identifier")
			if err != nil {
				return nil, err
			}
			args = append(args, arg)

			nextSep := p.advance()
			if nextSep.Type == tokentype.RIGHT_PAREN {
				break
			} else if nextSep.Type != tokentype.COMMA {
				return nil, loxerr.AtToken(nextSep, "Expected ')' after args.")
			}
		}
	}

	// After args is the function body starting with '{'.
	_, err = p.consume(tokentype.LEFT_BRACE, "Expected '{'.")
	if err != nil {
		return nil, err
	}

	// Parse the function body.
	funcBody, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	return &ast.FunctionStmt{
		Symbol: symbol,
		Args:   args,
		Body:   funcBody,
	}, nil
}

// Parse a var statement.
func (p *Parser) parseVarStatement() (*ast.VarStmt, error) {
	// LHS of the var declaration.
	lhsToken, err := p.consume(tokentype.IDENTIFIER, "Expected identifier after 'var'.")
	if err != nil {
		return nil, err
	}

	var rhs ast.Expr

	// Optional RHS of the declaration.
	if p.peek(1).Type == tokentype.EQUAL {
		p.advance()
		rhs, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}

	// Declaration is a statement that must be terminated with a semicolon.
	_, err = p.consume(tokentype.SEMICOLON, "Expected ';'.")
	if err != nil {
		return nil, err
	}

	return &ast.VarStmt{
		Left:  lhsToken,
		Right: rhs,
	}, nil
}

// Parse an expression.
func (p *Parser) parseExpression() (ast.Expr, error) {
	return p.parseAssignment()
}

// Parse an assignment expression.
func (p *Parser) parseAssignment() (ast.Expr, error) {
	expr, err := p.parseLogicalOr()
	if err != nil {
		return nil, err
	}

	if p.peek(1).Type == tokentype.EQUAL {
		equalsToken := p.peek(0)
		p.advance()
		right, err := p.parseAssignment()
		if err != nil {
			return nil, err
		}

		// Only allow assignment to a VarExpr.
		if varExpr, ok := expr.(*ast.VarExpr); ok {
			expr, err = &ast.AssignExpr{Left: varExpr.Name, Right: right}, nil
		} else {
			expr, err = nil, loxerr.AtToken(equalsToken, "Invalid assignment target.")
		}
	}
	return expr, err
}

func (p *Parser) parseLogicalOr() (ast.Expr, error) {
	expr, err := p.parseLogicalAnd()
	if err != nil {
		return nil, err
	}

	for p.peekMatches(1, tokentype.OR) {
		left := expr
		operator := p.advance()
		right, err := p.parseLogicalAnd()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) parseLogicalAnd() (ast.Expr, error) {
	expr, err := p.parseEquality()
	if err != nil {
		return nil, err
	}

	for p.peekMatches(1, tokentype.AND) {
		left := expr
		operator := p.advance()
		right, err := p.parseLogicalAnd()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return expr, nil
}

// Parse an equality expression.
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

// Parse a comparison expression.
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

// Parse a term expression.
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

// Parse a factor expression.
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

// Parse a unary expression.
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
	return p.parseCall()
}

// Parse a call expression.
func (p *Parser) parseCall() (ast.Expr, error) {
	// Leftmost part of the call should be a primary expr.
	expr, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}

	// Find all consecutive call operators.
	for nextToken := p.peek(1); nextToken.Type == tokentype.LEFT_PAREN; nextToken = p.peek(1) {
		p.advance()

		// Extract the args for this call, if any.
		args := []ast.Expr{}
		if p.peekMatches(1, tokentype.RIGHT_PAREN) {
			// No args, move on.
			p.advance()
		} else {
			// Parse the args.
			for {
				arg, err := p.parseExpression()
				if err != nil {
					return nil, err
				}
				args = append(args, arg)

				// After the arg, there must be either a comma or a closing parentheses.
				// If closing parentheses, done parsing the args.
				// Continue parsing args if it's a comma.
				nextSep := p.advance()
				if nextSep.Type == tokentype.RIGHT_PAREN {
					break
				} else if nextSep.Type != tokentype.COMMA {
					return nil, loxerr.AtToken(nextSep, "Expected ')' after call.")
				}
			}
		}

		// The current call becomes the callee of the next call.
		expr = &ast.CallExpr{
			Callee:    expr,
			OpenParen: nextToken,
			Args:      args,
		}
	}

	return expr, nil
}

// Parse a primary expression.
func (p *Parser) parsePrimary() (ast.Expr, error) {
	matchTokens := []tokentype.TokenType{
		tokentype.NUMBER, tokentype.STRING,
		tokentype.TRUE, tokentype.FALSE,
		tokentype.IDENTIFIER, tokentype.NIL,
	}
	if p.peekMatches(1, matchTokens...) {
		matched := p.advance()
		switch matched.Type {
		case tokentype.TRUE:
			return &ast.LiteralExpr{Value: true}, nil
		case tokentype.FALSE:
			return &ast.LiteralExpr{Value: false}, nil
		case tokentype.IDENTIFIER:
			return &ast.VarExpr{Name: matched}, nil
		default:
			return &ast.LiteralExpr{Value: matched.Literal}, nil
		}
	} else if p.peekMatches(1, tokentype.LEFT_PAREN) {
		p.advance()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		_, err = p.consume(tokentype.RIGHT_PAREN, "Expected ')' after expression.")
		if err != nil {
			return nil, err
		}
		return &ast.GroupingExpr{Expression: expr}, nil
	} else {
		return nil, loxerr.AtToken(p.peek(1), "Expected expression.")
	}
}

// Advance the current pointer to the next token if it matches typeToMatch.
func (p *Parser) consume(typeToMatch tokentype.TokenType, errorMessage string) (*token.Token, error) {
	nextToken := p.peek(1)
	if nextToken.Type == typeToMatch {
		p.current++
		return nextToken, nil
	}
	return nil, loxerr.AtToken(nextToken, errorMessage)
}

// Advance the current pointer to the next token.
func (p *Parser) advance() *token.Token {
	result := p.peek(1)
	if !p.isAtEnd() {
		p.current++
	}
	return result
}

// Whether or not the token that is lookahead in front of the current token matches
// any of the provided TokenType items.
func (p *Parser) peekMatches(lookahead int, tokenTypes ...tokentype.TokenType) bool {
	for i := 0; i < len(tokenTypes); i++ {
		if p.peek(lookahead).Type == tokenTypes[i] {
			return true
		}
	}
	return false
}

// Get the token that is lookahead in front of the current token.
func (p *Parser) peek(lookahead int) *token.Token {
	return p.tokens[p.current+lookahead-1]
}

// Whether the parser is at the end of input.
func (p *Parser) isAtEnd() bool {
	return p.peek(1).Type == tokentype.EOF
}

// Move the token pointer to the beginning of the next statement.
// This should be called after handling an error.
func (p *Parser) synchronize() {
	stopAtTokens := []tokentype.TokenType{
		tokentype.CLASS, tokentype.FUN, tokentype.VAR,
		tokentype.FOR, tokentype.IF, tokentype.WHILE,
		tokentype.PRINT, tokentype.RETURN,
	}

	for !p.isAtEnd() {
		p.current++
		if p.peekMatches(0, tokentype.SEMICOLON) || p.peekMatches(1, stopAtTokens...) {
			return
		}
	}
}
