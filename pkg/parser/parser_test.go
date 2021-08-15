package parser

import (
	"testing"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
	"github.com/stretchr/testify/assert"
)

func strToken(val string) *token.Token {
	return &token.Token{
		Type:    tokentype.STRING,
		Lexeme:  val,
		Literal: val,
		Line:    1,
	}
}

func symToken(tokenType tokentype.TokenType, lexeme string) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: nil,
		Line:    1,
	}
}

func eofToken() *token.Token {
	return &token.Token{
		Type:    tokentype.EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    1,
	}
}

func printToken() *token.Token {
	return &token.Token{
		Type:    tokentype.PRINT,
		Lexeme:  "print",
		Literal: nil,
		Line:    1,
	}
}

func boolToken(lexeme string) *token.Token {
	var tokenType tokentype.TokenType
	if lexeme == "true" {
		tokenType = tokentype.TRUE
	} else if lexeme == "false" {
		tokenType = tokentype.FALSE
	} else {
		panic("Invalid value in test.")
	}

	return &token.Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: nil,
		Line:    1,
	}
}

func assertTokensEqual(t *testing.T, a *token.Token, b *token.Token) {
	assert.Equal(t, a.Literal, b.Literal)
	assert.Equal(t, a.Lexeme, b.Lexeme)
}

func assertIsBinaryExpr(t *testing.T, expr ast.Expr) *ast.BinaryExpr {
	tree, ok := expr.(*ast.BinaryExpr)
	assert.True(t, ok)
	return tree
}

func assertIsUnaryExpr(t *testing.T, expr ast.Expr) *ast.UnaryExpr {
	tree, ok := expr.(*ast.UnaryExpr)
	assert.True(t, ok)
	return tree
}

func assertIsGroupingExpr(t *testing.T, expr ast.Expr) *ast.GroupingExpr {
	tree, ok := expr.(*ast.GroupingExpr)
	assert.True(t, ok)
	return tree
}

func assertIsLiteralExpr(t *testing.T, expr ast.Expr) *ast.LiteralExpr {
	tree, ok := expr.(*ast.LiteralExpr)
	assert.True(t, ok)
	return tree
}

func assertIsPrintStmt(t *testing.T, stmt ast.Stmt) *ast.PrintStmt {
	tree, ok := stmt.(*ast.PrintStmt)
	assert.True(t, ok)
	return tree
}

func assertIsExprStmt(t *testing.T, stmt ast.Stmt) *ast.ExprStmt {
	tree, ok := stmt.(*ast.ExprStmt)
	assert.True(t, ok)
	return tree
}

func assertBinaryExprOfLiterals(t *testing.T, actual *ast.BinaryExpr, expectedLhs interface{}, expectedOp *token.Token, expectedRhs interface{}) {
	lhs, ok := actual.Left.(*ast.LiteralExpr)
	assert.True(t, ok)
	assert.Equal(t, lhs.Value, expectedLhs)

	op := actual.Operator
	assertTokensEqual(t, expectedOp, op)

	rhs, ok := actual.Right.(*ast.LiteralExpr)
	assert.True(t, ok)
	assert.Equal(t, rhs.Value, expectedRhs)
}

func assertUnaryExpressionOfLiteral(t *testing.T, actual *ast.UnaryExpr, expectedOp *token.Token, expectedRhs interface{}) {
	assertTokensEqual(t, actual.Operator, expectedOp)

	rhs, ok := actual.Right.(*ast.LiteralExpr)
	assert.True(t, ok)
	assert.Equal(t, rhs.Value, expectedRhs)
}

func testBinaryExpressionWithLiterals(t *testing.T, expectedOp *token.Token) {
	lhsValue := "lhsToken"
	rhsValue := "rhsToken"

	// "lhsToken" <expectedOp> "rhsToken" <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseExpression()
	assert.Nil(t, err)

	expr := assertIsBinaryExpr(t, tree)
	assertBinaryExprOfLiterals(t, expr, lhsValue, expectedOp, rhsValue)
}

func testUnaryExpression(t *testing.T, expectedOp *token.Token) {
	rhsValue := "rhsToken"

	// <expectedOp> "lhsToken" <EOF>
	parser := NewParser([]*token.Token{
		expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseExpression()
	assert.Nil(t, err)

	expr := assertIsUnaryExpr(t, tree)
	assertUnaryExpressionOfLiteral(t, expr, expectedOp, rhsValue)
}

func TestParse_MissingEOF(t *testing.T) {
	// print "lhsToken" +"rhsToken";
	parser := NewParser([]*token.Token{
		printToken(), strToken("lhsVaue"), symToken(tokentype.PLUS, "+"), strToken("rhsValue"),
		symToken(tokentype.SEMICOLON, ";"),
	})
	program, errs := parser.Parse()
	assert.Nil(t, program)
	assert.NotEmpty(t, errs)
}

func TestParse_EmptyMissingEOF(t *testing.T) {
	parser := NewParser([]*token.Token{})
	program, errs := parser.Parse()
	assert.Nil(t, program)
	assert.NotEmpty(t, errs)
}

func TestParse_OnlyEOF(t *testing.T) {
	// <EOF>
	parser := NewParser([]*token.Token{eofToken()})
	program, errs := parser.Parse()
	assert.Empty(t, errs)
	assert.Empty(t, program.Statements)
}

func TestParseExpression_Equality(t *testing.T) {
	testBinaryExpressionWithLiterals(t, symToken(tokentype.BANG_EQUAL, "!="))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.EQUAL_EQUAL, "=="))
}

func TestParseExpression_Comparison(t *testing.T) {
	testBinaryExpressionWithLiterals(t, symToken(tokentype.GREATER, ">"))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.GREATER_EQUAL, ">="))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.LESS, "<"))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.LESS_EQUAL, "<="))
}

func TestParseExpression_Term(t *testing.T) {
	testBinaryExpressionWithLiterals(t, symToken(tokentype.MINUS, "-"))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.PLUS, "+"))
}

func TestParseExpression_Factor(t *testing.T) {
	testBinaryExpressionWithLiterals(t, symToken(tokentype.SLASH, "/"))
	testBinaryExpressionWithLiterals(t, symToken(tokentype.STAR, "*"))
}

func TestParseExpression_Unary(t *testing.T) {
	testUnaryExpression(t, symToken(tokentype.BANG, "!"))
	testUnaryExpression(t, symToken(tokentype.MINUS, "-"))
}

func TestParseExpression_Primary_AsUnaryInParentheses(t *testing.T) {
	expectedOp := symToken(tokentype.BANG, "!")
	rhsValue := "rhsToken"

	// <expectedOp> "lhsToken" <EOF>
	parser := NewParser([]*token.Token{
		symToken(tokentype.LEFT_PAREN, "("), expectedOp, strToken(rhsValue),
		symToken(tokentype.RIGHT_PAREN, ")"), eofToken(),
	})
	tree, err := parser.parseExpression()
	assert.Nil(t, err)

	groupExpr := assertIsGroupingExpr(t, tree)
	unaryExpr := assertIsUnaryExpr(t, groupExpr.Expression)
	assertUnaryExpressionOfLiteral(t, unaryExpr, expectedOp, rhsValue)
}

func TestParseExpression_Primary_MissingClosingParentheses(t *testing.T) {
	expectedOp := symToken(tokentype.BANG, "!")
	rhsValue := "rhsToken"

	// <expectedOp> "lhsToken" <EOF>
	parser := NewParser([]*token.Token{
		symToken(tokentype.LEFT_PAREN, "("), expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseExpression()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseExpression_Primary_BooleanLiterals(t *testing.T) {
	var parser *Parser
	var tree ast.Expr
	var err error
	var expr *ast.LiteralExpr

	parser = NewParser([]*token.Token{
		boolToken("true"), eofToken(),
	})
	tree, err = parser.parseExpression()
	assert.Nil(t, err)
	expr = assertIsLiteralExpr(t, tree)
	assert.Equal(t, expr.Value, true)

	parser = NewParser([]*token.Token{
		boolToken("false"), eofToken(),
	})
	tree, err = parser.parseExpression()
	assert.Nil(t, err)
	expr = assertIsLiteralExpr(t, tree)
	assert.Equal(t, expr.Value, false)
}

func TestParseExpression_MissingRhs(t *testing.T) {
	lhsValue := "lhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// "lhsToken" + <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, eofToken(),
	})
	tree, err := parser.parseExpression()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParsePrintStatement_Basic(t *testing.T) {
	lhsValue := "lhsToken"
	rhsValue := "rhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" +"rhsToken"; <EOF>
	parser := NewParser([]*token.Token{
		printToken(), strToken(lhsValue), expectedOp, strToken(rhsValue),
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	stmt := assertIsPrintStmt(t, tree)
	expr := assertIsBinaryExpr(t, stmt.Expr)
	assertBinaryExprOfLiterals(t, expr, lhsValue, expectedOp, rhsValue)
}

func TestParsePrintStatement_BadExpression(t *testing.T) {
	lhsValue := "lhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" + <EOF>
	parser := NewParser([]*token.Token{
		printToken(), strToken(lhsValue), expectedOp,
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParsePrintStatement_MissingSemicolon(t *testing.T) {
	lhsValue := "lhsToken"
	rhsValue := "rhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" + "rhsToken" <EOF>
	parser := NewParser([]*token.Token{
		printToken(), strToken(lhsValue), expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseExpressionStmt_Basic(t *testing.T) {
	lhsValue := "lhsToken"
	rhsValue := "rhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" +"rhsToken"; <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, strToken(rhsValue),
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	stmt := assertIsExprStmt(t, tree)
	expr := assertIsBinaryExpr(t, stmt.Expr)
	assertBinaryExprOfLiterals(t, expr, lhsValue, expectedOp, rhsValue)
}

func TestParseExpressionStmt_BadExpression(t *testing.T) {
	lhsValue := "lhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" + <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseExpressionStmt_MissingSemicolon(t *testing.T) {
	lhsValue := "lhsToken"
	rhsValue := "rhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// print "lhsToken" + <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}
