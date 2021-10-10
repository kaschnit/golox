package parser

import (
	"strconv"
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

func numToken(val int) *token.Token {
	return &token.Token{
		Type:    tokentype.NUMBER,
		Lexeme:  strconv.Itoa(val),
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
	return symToken(tokentype.EOF, "")
}

func printToken() *token.Token {
	return symToken(tokentype.PRINT, "print")
}

func ifToken() *token.Token {
	return symToken(tokentype.IF, "if")
}

func elseToken() *token.Token {
	return symToken(tokentype.ELSE, "else")
}

func varToken() *token.Token {
	return symToken(tokentype.VAR, "var")
}

func whileToken() *token.Token {
	return symToken(tokentype.WHILE, "while")
}

func forToken() *token.Token {
	return symToken(tokentype.FOR, "for")
}

func boolToken(val bool) *token.Token {
	var tokenType tokentype.TokenType
	var lexeme string
	if val {
		lexeme = "true"
		tokenType = tokentype.TRUE
	} else {
		lexeme = "false"
		tokenType = tokentype.FALSE
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

func assertIsIfStmt(t *testing.T, stmt ast.Stmt) *ast.IfStmt {
	tree, ok := stmt.(*ast.IfStmt)
	assert.True(t, ok)
	return tree
}

func assertIsWhileStmt(t *testing.T, stmt ast.Stmt) *ast.WhileStmt {
	tree, ok := stmt.(*ast.WhileStmt)
	assert.True(t, ok)
	return tree
}

func assertIsVarStmt(t *testing.T, stmt ast.Stmt) *ast.VarStmt {
	tree, ok := stmt.(*ast.VarStmt)
	assert.True(t, ok)
	return tree
}

func assertIsExprStmt(t *testing.T, stmt ast.Stmt) *ast.ExprStmt {
	tree, ok := stmt.(*ast.ExprStmt)
	assert.True(t, ok)
	return tree
}

func assertIsBlockStmt(t *testing.T, stmt ast.Stmt) *ast.BlockStmt {
	tree, ok := stmt.(*ast.BlockStmt)
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
		boolToken(true), eofToken(),
	})
	tree, err = parser.parseExpression()
	assert.Nil(t, err)
	expr = assertIsLiteralExpr(t, tree)
	assert.Equal(t, expr.Value, true)

	parser = NewParser([]*token.Token{
		boolToken(false), eofToken(),
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
	expr := assertIsBinaryExpr(t, stmt.Expression)
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

	// "lhsToken" + "rhsToken" ; <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, strToken(rhsValue),
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	stmt := assertIsExprStmt(t, tree)
	expr := assertIsBinaryExpr(t, stmt.Expression)
	assertBinaryExprOfLiterals(t, expr, lhsValue, expectedOp, rhsValue)
}

func TestParseExpressionStmt_BadExpression(t *testing.T) {
	lhsValue := "lhsToken"
	expectedOp := symToken(tokentype.PLUS, "+")

	// "lhsToken" + ; <EOF>
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

	// "lhsToken" + "rhsToken" <EOF>
	parser := NewParser([]*token.Token{
		strToken(lhsValue), expectedOp, strToken(rhsValue), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseExpressionStmt_EmptyStmt(t *testing.T) {
	// ; <EOF>
	parser := NewParser([]*token.Token{
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	stmt := assertIsExprStmt(t, tree)
	expr := assertIsLiteralExpr(t, stmt.Expression)
	assert.Nil(t, expr.Value)
}

func TestParseVarStmt_Basic(t *testing.T) {
	lhsName := "a"
	rhsVal := "hello"

	// var a = "hello";
	parser := NewParser([]*token.Token{
		varToken(), symToken(tokentype.IDENTIFIER, lhsName), symToken(tokentype.EQUAL, "="),
		strToken(rhsVal), symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	varStmt := assertIsVarStmt(t, tree)
	rhsExpr := assertIsLiteralExpr(t, varStmt.Right)
	assert.Equal(t, tokentype.IDENTIFIER, varStmt.Left.Type)
	assert.Equal(t, lhsName, varStmt.Left.Lexeme)
	assert.Equal(t, rhsVal, rhsExpr.Value)
}

func TestParseVarStmt_InvalidLhsNumerical(t *testing.T) {
	lhsVal := 1
	rhsVal := "hello"

	// var 1 = "hello";
	parser := NewParser([]*token.Token{
		varToken(), numToken(lhsVal), symToken(tokentype.EQUAL, "="), strToken(rhsVal),
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseVarStmt_InvalidLhsString(t *testing.T) {
	lhsVal := "hello"
	rhsVal := 99

	// var "hello" = 99;
	parser := NewParser([]*token.Token{
		varToken(), strToken(lhsVal), symToken(tokentype.EQUAL, "="), numToken(rhsVal),
		symToken(tokentype.SEMICOLON, ";"), eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, tree)
	assert.Error(t, err)
}

func TestParseIfStmt_NoBlockNoElse(t *testing.T) {
	ifBranchPrint := "if branch"

	// if (true) print "if branch"; <EOF>
	parser := NewParser([]*token.Token{
		ifToken(), symToken(tokentype.LEFT_PAREN, "("), boolToken(true), symToken(tokentype.RIGHT_PAREN, ")"),
		printToken(), strToken(ifBranchPrint), symToken(tokentype.SEMICOLON, ";"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	ifStmt := assertIsIfStmt(t, tree)
	cond := assertIsLiteralExpr(t, ifStmt.Condition)
	printStmt := assertIsPrintStmt(t, ifStmt.ThenStatement)
	printStmtExpr := assertIsLiteralExpr(t, printStmt.Expression)
	assert.Equal(t, cond.Value, true)
	assert.Equal(t, printStmtExpr.Value, ifBranchPrint)
	assert.Nil(t, ifStmt.ElseStatement)
}

func TestParseIfStmt_NoBlockWithElse(t *testing.T) {
	ifBranchPrint := "if branch"
	elseBranchPrint := "else branch"

	// if (true) print "if branch"; else print "else branch"; <EOF>
	parser := NewParser([]*token.Token{
		ifToken(), symToken(tokentype.LEFT_PAREN, "("), boolToken(true), symToken(tokentype.RIGHT_PAREN, ")"),
		printToken(), strToken(ifBranchPrint), symToken(tokentype.SEMICOLON, ";"),
		elseToken(), printToken(), strToken(elseBranchPrint), symToken(tokentype.SEMICOLON, ";"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	ifStmt := assertIsIfStmt(t, tree)
	cond := assertIsLiteralExpr(t, ifStmt.Condition)
	ifPrintStmt := assertIsPrintStmt(t, ifStmt.ThenStatement)
	ifPrintStmtExpr := assertIsLiteralExpr(t, ifPrintStmt.Expression)
	assert.Equal(t, cond.Value, true)
	assert.Equal(t, ifPrintStmtExpr.Value, ifBranchPrint)

	elsePrintStmt := assertIsPrintStmt(t, ifStmt.ElseStatement)
	elsePrintStmtExpr := assertIsLiteralExpr(t, elsePrintStmt.Expression)
	assert.Equal(t, cond.Value, true)
	assert.Equal(t, elsePrintStmtExpr.Value, elseBranchPrint)
}

func TestParseWhileStmt_Basic(t *testing.T) {
	bodyPrintStmt := "while loop printing"
	bodyExprStmtLhs := "1"
	bodyExprStmtRhs := "2"
	bodyExprStmtOp := symToken(tokentype.PLUS, "+")

	// while (false) { print "while loop printing"; 1 + 2; } <EOF>
	parser := NewParser([]*token.Token{
		whileToken(), symToken(tokentype.LEFT_PAREN, "("), boolToken(false), symToken(tokentype.RIGHT_PAREN, ")"),
		symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		strToken(bodyExprStmtLhs), bodyExprStmtOp, strToken(bodyExprStmtRhs), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	whileStmt := assertIsWhileStmt(t, tree)
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	block := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Equal(t, cond.Value, false)
	assert.Len(t, block.Statements, 2)

	printStmt := assertIsPrintStmt(t, block.Statements[0])
	printStmtExpr := assertIsLiteralExpr(t, printStmt.Expression)
	assert.Equal(t, printStmtExpr.Value, bodyPrintStmt)

	exprStmt := assertIsExprStmt(t, block.Statements[1])
	exprStmtExpr := assertIsBinaryExpr(t, exprStmt.Expression)
	assertBinaryExprOfLiterals(t, exprStmtExpr, bodyExprStmtLhs, bodyExprStmtOp, bodyExprStmtRhs)
}

func TestParseForStmt_Empty(t *testing.T) {
	bodyPrintStmt := "for loop printing"

	// for (;;) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("),
		symToken(tokentype.SEMICOLON, ";"), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_PAREN, ")"), symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop desugared to a while loop.
	whileStmt := assertIsWhileStmt(t, tree)

	// Condition is implicitly true if not provided.
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	assert.Equal(t, cond.Value, true)

	// Block has exactly 1 statement, the print statement.
	block := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, block.Statements, 1)
	assertIsPrintStmt(t, block.Statements[0])
}

func TestParseForStmt_OnlyInitialiazer_VarStmt(t *testing.T) {
	bodyPrintStmt := "for loop printing"
	varName := "x"
	varValue := 34

	// for (var x = 34;;) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("),
		varToken(), symToken(tokentype.IDENTIFIER, varName),
		symToken(tokentype.EQUAL, "="), numToken(varValue), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.SEMICOLON, ";"), symToken(tokentype.RIGHT_PAREN, ")"),
		symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop with initializer is desugared to an initializer + while loop within a block.
	blockWrapperStmt := assertIsBlockStmt(t, tree)
	assert.Len(t, blockWrapperStmt.Statements, 2)
	initializerStmt := assertIsVarStmt(t, blockWrapperStmt.Statements[0])
	initRight := assertIsLiteralExpr(t, initializerStmt.Right)
	assert.Equal(t, varName, initializerStmt.Left.Lexeme)
	assert.Equal(t, varValue, initRight.Value)

	whileStmt := assertIsWhileStmt(t, blockWrapperStmt.Statements[1])

	// Condition is implicitly true if not provided.
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	assert.Equal(t, cond.Value, true)

	// Block has exactly 1 statement, the print statement.
	block := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, block.Statements, 1)
	assertIsPrintStmt(t, block.Statements[0])
}

func TestParseForStmt_OnlyInitialiazer_ExprStmt(t *testing.T) {
	bodyPrintStmt := "for loop printing"
	initExprValue := "hello"

	// for ("hello";;) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("), strToken(initExprValue),
		symToken(tokentype.SEMICOLON, ";"), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_PAREN, ")"), symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop with initializer is desugared to an initializer + while loop within a block.
	blockWrapperStmt := assertIsBlockStmt(t, tree)
	assert.Len(t, blockWrapperStmt.Statements, 2)
	initializerStmt := assertIsExprStmt(t, blockWrapperStmt.Statements[0])
	initializeExpr := assertIsLiteralExpr(t, initializerStmt.Expression)
	assert.Equal(t, initExprValue, initializeExpr.Value)

	whileStmt := assertIsWhileStmt(t, blockWrapperStmt.Statements[1])

	// Condition is implicitly true if not provided.
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	assert.Equal(t, cond.Value, true)

	// Block has exactly 1 statement, the print statement.
	block := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, block.Statements, 1)
	assertIsPrintStmt(t, block.Statements[0])
}

func TestParseForStmt_OnlyCondition(t *testing.T) {
	bodyPrintStmt := "for loop printing"
	condLeftVal := 1
	condRightVal := 2

	// for (; 1 < 2;) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("), symToken(tokentype.SEMICOLON, ";"),
		numToken(condLeftVal), symToken(tokentype.LESS, "<"), numToken(condRightVal),
		symToken(tokentype.SEMICOLON, ";"), symToken(tokentype.RIGHT_PAREN, ")"), symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop desugared to a while loop.
	whileStmt := assertIsWhileStmt(t, tree)

	// Condition provided.
	cond := assertIsBinaryExpr(t, whileStmt.Condition)
	condLeft := assertIsLiteralExpr(t, cond.Left)
	condRight := assertIsLiteralExpr(t, cond.Right)
	assert.Equal(t, condLeftVal, condLeft.Value)
	assert.Equal(t, tokentype.LESS, cond.Operator.Type)
	assert.Equal(t, condRightVal, condRight.Value)

	// Block has exactly 1 statement, the print statement.
	block := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, block.Statements, 1)
	assertIsPrintStmt(t, block.Statements[0])
}

func TestParseForStmt_OnlyIncrement(t *testing.T) {
	bodyPrintStmt := "for loop printing"
	incrLeftVal := 1
	incrRightVal := 987

	// for (;; 1 + 987) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.SEMICOLON, ";"), numToken(incrLeftVal),
		symToken(tokentype.PLUS, "+"), numToken(incrRightVal),
		symToken(tokentype.RIGHT_PAREN, ")"), symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop desugared to a while loop.
	whileStmt := assertIsWhileStmt(t, tree)

	// Condition is implicitly true if not provided.
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	assert.Equal(t, cond.Value, true)

	// Block has a block within it.
	// The outer block contains the inner block, then the increment.
	// The inner block contains the actual loop body.
	outerBlock := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, outerBlock.Statements, 2)
	innerBlock := assertIsBlockStmt(t, outerBlock.Statements[0])
	assert.Len(t, innerBlock.Statements, 1)
	increment := assertIsExprStmt(t, outerBlock.Statements[1])
	incrementExpr := assertIsBinaryExpr(t, increment.Expression)
	incrementLeft := assertIsLiteralExpr(t, incrementExpr.Left)
	assert.Equal(t, incrLeftVal, incrementLeft.Value)
	incrementRight := assertIsLiteralExpr(t, incrementExpr.Right)
	assert.Equal(t, incrRightVal, incrementRight.Value)
}

func TestParseForStmt_AllLoopExprs(t *testing.T) {
	bodyPrintStmt := "for loop printing"
	initVarName := "y"
	initVarValue := "abc"
	condValue := false
	incrLeftVal := 1
	incrRightVal := 987

	// for (var y = "abc"; false; 1 + 987) { print "for loop printing"; } <EOF>
	parser := NewParser([]*token.Token{
		forToken(), symToken(tokentype.LEFT_PAREN, "("),
		varToken(), symToken(tokentype.IDENTIFIER, initVarName), symToken(tokentype.EQUAL, "="),
		strToken(initVarValue), symToken(tokentype.SEMICOLON, ";"), boolToken(condValue),
		symToken(tokentype.SEMICOLON, ";"), numToken(incrLeftVal),
		symToken(tokentype.PLUS, "+"), numToken(incrRightVal),
		symToken(tokentype.RIGHT_PAREN, ")"), symToken(tokentype.LEFT_BRACE, "{"),
		printToken(), strToken(bodyPrintStmt), symToken(tokentype.SEMICOLON, ";"),
		symToken(tokentype.RIGHT_BRACE, "}"),
		eofToken(),
	})
	tree, err := parser.parseStatement()
	assert.Nil(t, err)

	// For loop with initializer is desugared to an initializer + while loop within a block.
	blockWrapperStmt := assertIsBlockStmt(t, tree)
	assert.Len(t, blockWrapperStmt.Statements, 2)
	initVarStmt := assertIsVarStmt(t, blockWrapperStmt.Statements[0])
	initVarStmtRight := assertIsLiteralExpr(t, initVarStmt.Right)
	assert.Equal(t, initVarName, initVarStmt.Left.Lexeme)
	assert.Equal(t, initVarValue, initVarStmtRight.Value)

	// For loop desugared to a while loop.
	whileStmt := assertIsWhileStmt(t, blockWrapperStmt.Statements[1])

	// Condition provided.
	cond := assertIsLiteralExpr(t, whileStmt.Condition)
	assert.Equal(t, false, cond.Value)

	// Block has a block within it.
	// The outer block contains the inner block, then the increment.
	// The inner block contains the actual loop body.
	outerBlock := assertIsBlockStmt(t, whileStmt.LoopStatement)
	assert.Len(t, outerBlock.Statements, 2)
	innerBlock := assertIsBlockStmt(t, outerBlock.Statements[0])
	assert.Len(t, innerBlock.Statements, 1)
	increment := assertIsExprStmt(t, outerBlock.Statements[1])
	incrementExpr := assertIsBinaryExpr(t, increment.Expression)
	incrementLeft := assertIsLiteralExpr(t, incrementExpr.Left)
	assert.Equal(t, incrLeftVal, incrementLeft.Value)
	incrementRight := assertIsLiteralExpr(t, incrementExpr.Right)
	assert.Equal(t, incrRightVal, incrementRight.Value)
}
