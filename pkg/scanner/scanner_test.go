package scanner

import (
	"fmt"
	"testing"

	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
	"github.com/stretchr/testify/assert"
)

const floatDelta = 0.0001

func verifyNextScanTokenIsEOF(t *testing.T, scanner *Scanner, expectedLine int) {
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.EOF, token.Type)
	assert.Equal(t, "", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, expectedLine, token.Line)
	verifyNextScanTokenIsPastEOFError(t, scanner)
	verifyNextScanTokenIsPastEOFError(t, scanner)
}

func verifyNextScanTokenIsPastEOFError(t *testing.T, scanner *Scanner) {
	token, err := scanner.ScanToken()
	assert.Nil(t, token)
	assert.Error(t, err)
}

func verifyScanTokenSingleIdentifier(t *testing.T, input string) {
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, input, token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)
}

func verifyScanTokenSingleKeyword(t *testing.T, input string, expectedType tokentype.TokenType) {
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, expectedType, token.Type)
	assert.Equal(t, input, token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)
}

func verifyScanTokenSingleFloat(t *testing.T, input string, expectedLiteral float64) {
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.NUMBER, token.Type)
	assert.Equal(t, input, token.Lexeme)
	assert.InDelta(t, expectedLiteral, token.Literal, floatDelta)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)
}

func verifyScanTokenSingle(t *testing.T, input string, expectedType tokentype.TokenType, expectedLiteral interface{}) {
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, expectedType, token.Type)
	assert.Equal(t, input, token.Lexeme)
	assert.Equal(t, expectedLiteral, token.Literal, floatDelta)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)
}

func TestScanTokenInteger_Success(t *testing.T) {
	verifyScanTokenSingleFloat(t, "123", 123.0)
	verifyScanTokenSingleFloat(t, "0", 0.0)
	verifyScanTokenSingleFloat(t, "1", 1.0)
}

func TestScanTokenFloat_Success(t *testing.T) {
	verifyScanTokenSingleFloat(t, "123.5050", 123.505)
	verifyScanTokenSingleFloat(t, "987.1", 987.1)
	verifyScanTokenSingleFloat(t, "3.0", 3.0)
}

func TestScanTokenKeyword_Success(t *testing.T) {
	verifyScanTokenSingleKeyword(t, "and", tokentype.AND)
	verifyScanTokenSingleKeyword(t, "class", tokentype.CLASS)
	verifyScanTokenSingleKeyword(t, "else", tokentype.ELSE)
	verifyScanTokenSingleKeyword(t, "false", tokentype.FALSE)
	verifyScanTokenSingleKeyword(t, "fun", tokentype.FUN)
	verifyScanTokenSingleKeyword(t, "for", tokentype.FOR)
	verifyScanTokenSingleKeyword(t, "if", tokentype.IF)
	verifyScanTokenSingleKeyword(t, "nil", tokentype.NIL)
	verifyScanTokenSingleKeyword(t, "or", tokentype.OR)
	verifyScanTokenSingleKeyword(t, "print", tokentype.PRINT)
	verifyScanTokenSingleKeyword(t, "return", tokentype.RETURN)
	verifyScanTokenSingleKeyword(t, "super", tokentype.SUPER)
	verifyScanTokenSingleKeyword(t, "this", tokentype.THIS)
	verifyScanTokenSingleKeyword(t, "true", tokentype.TRUE)
	verifyScanTokenSingleKeyword(t, "var", tokentype.VAR)
	verifyScanTokenSingleKeyword(t, "while", tokentype.WHILE)
}

func TestScanTokenString_BasicString(t *testing.T) {
	expected := "hello"
	input := fmt.Sprintf(`"%s"`, expected)
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.STRING, token.Type)
	assert.Equal(t, expected, token.Lexeme)
	assert.Equal(t, expected, token.Literal)
	assert.Equal(t, token.Line, 1)
}

func TestScanTokenString_StringWithWhitespace(t *testing.T) {
	expected := " hello  \t"
	input := fmt.Sprintf(`"%s"`, expected)
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.STRING, token.Type)
	assert.Equal(t, expected, token.Lexeme)
	assert.Equal(t, expected, token.Literal)
	assert.Equal(t, token.Line, 1)
}
func TestScanTokenString_StringWithComment(t *testing.T) {
	expected := "hello // here's a comment"
	input := fmt.Sprintf(`"%s"`, expected)
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.STRING, token.Type)
	assert.Equal(t, expected, token.Lexeme)
	assert.Equal(t, expected, token.Literal)
	assert.Equal(t, token.Line, 1)
}

func TestScanTokenString_UnterminatedString(t *testing.T) {
	input := `"hello`
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Error(t, err)
	assert.Nil(t, token)
}
func TestScanTokenString_Multiline(t *testing.T) {
	expected := `hello
	
	a

	`
	input := fmt.Sprintf(`"%s"
	
`, expected)
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.STRING, token.Type)
	assert.Equal(t, expected, token.Lexeme)
	assert.Equal(t, expected, token.Literal)
	assert.Equal(t, token.Line, 5)
}

func TestScanTokenString_EmptyString(t *testing.T) {
	expected := ""
	input := fmt.Sprintf(`"%s"`, expected)
	scanner := NewScanner(input)
	token, err := scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.STRING, token.Type)
	assert.Equal(t, expected, token.Lexeme)
	assert.Equal(t, expected, token.Literal)
	assert.Equal(t, token.Line, 1)
}

func TestScanTokenIdentifier_KeywordCaseSensitive(t *testing.T) {
	verifyScanTokenSingleIdentifier(t, "AND")
	verifyScanTokenSingleIdentifier(t, "Else")
	verifyScanTokenSingleIdentifier(t, "fOr")
	verifyScanTokenSingleIdentifier(t, "IF")
	verifyScanTokenSingleIdentifier(t, "OR")
	verifyScanTokenSingleIdentifier(t, "SuPER")
	verifyScanTokenSingleIdentifier(t, "VAR")
}

func TestScanSymbols_Success(t *testing.T) {
	verifyScanTokenSingle(t, "(", tokentype.LEFT_PAREN, nil)
	verifyScanTokenSingle(t, ")", tokentype.RIGHT_PAREN, nil)
	verifyScanTokenSingle(t, "{", tokentype.LEFT_BRACE, nil)
	verifyScanTokenSingle(t, "}", tokentype.RIGHT_BRACE, nil)
	verifyScanTokenSingle(t, ",", tokentype.COMMA, nil)
	verifyScanTokenSingle(t, ".", tokentype.DOT, nil)
	verifyScanTokenSingle(t, "-", tokentype.MINUS, nil)
	verifyScanTokenSingle(t, "+", tokentype.PLUS, nil)
	verifyScanTokenSingle(t, ";", tokentype.SEMICOLON, nil)
	verifyScanTokenSingle(t, "/", tokentype.SLASH, nil)
	verifyScanTokenSingle(t, "*", tokentype.STAR, nil)
	verifyScanTokenSingle(t, "!", tokentype.BANG, nil)
	verifyScanTokenSingle(t, "!=", tokentype.BANG_EQUAL, nil)
	verifyScanTokenSingle(t, "=", tokentype.EQUAL, nil)
	verifyScanTokenSingle(t, "==", tokentype.EQUAL_EQUAL, nil)
	verifyScanTokenSingle(t, "<", tokentype.LESS, nil)
	verifyScanTokenSingle(t, "<=", tokentype.LESS_EQUAL, nil)
	verifyScanTokenSingle(t, ">", tokentype.GREATER, nil)
	verifyScanTokenSingle(t, ">=", tokentype.GREATER_EQUAL, nil)
}

func TestScanTokenIdentifier_Success(t *testing.T) {
	verifyScanTokenSingleIdentifier(t, "helloworld")
	verifyScanTokenSingleIdentifier(t, "_helloworld")
	verifyScanTokenSingleIdentifier(t, "_123")
	verifyScanTokenSingleIdentifier(t, "_123")
	verifyScanTokenSingleIdentifier(t, "_123abc")
	verifyScanTokenSingleIdentifier(t, "_123_abc")
	verifyScanTokenSingleIdentifier(t, "hello123")
	verifyScanTokenSingleIdentifier(t, "hello_123_")
	verifyScanTokenSingleIdentifier(t, "hello123_")
}

func TestScanToken_CountsLines(t *testing.T) {
	var token *token.Token
	var err error

	input := `some_ident other_ident
	hello
123
		456
`
	scanner := NewScanner(input)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "some_ident", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "other_ident", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "hello", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 2, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.NUMBER, token.Type)
	assert.Equal(t, "123", token.Lexeme)
	assert.InDelta(t, 123.0, token.Literal, floatDelta)
	assert.Equal(t, 3, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.NUMBER, token.Type)
	assert.Equal(t, "456", token.Lexeme)
	assert.InDelta(t, 456.0, token.Literal, floatDelta)
	assert.Equal(t, 4, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.EOF, token.Type)
	assert.Equal(t, "", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 5, token.Line)
}

func TestScanToken_SkipsWhitespace(t *testing.T) {
	var token *token.Token
	var err error

	input := "abc \t def       for 1 true    \t "
	scanner := NewScanner(input)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "abc", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "def", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.FOR, token.Type)
	assert.Equal(t, "for", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.NUMBER, token.Type)
	assert.Equal(t, "1", token.Lexeme)
	assert.InDelta(t, 1.0, token.Literal, floatDelta)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.TRUE, token.Type)
	assert.Equal(t, "true", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.EOF, token.Type)
	assert.Equal(t, "", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
}

func TestScanToken_InvalidCharacterFailsThenRecovers(t *testing.T) {
	var token *token.Token
	var err error

	input := "^ return ^^^ 123" // an invalid character
	scanner := NewScanner(input)

	token, err = scanner.ScanToken()
	assert.Nil(t, token)
	assert.Error(t, err)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.RETURN, token.Type)
	assert.Equal(t, "return", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, token)
	assert.Error(t, err)
	token, err = scanner.ScanToken()
	assert.Nil(t, token)
	assert.Error(t, err)
	token, err = scanner.ScanToken()
	assert.Nil(t, token)
	assert.Error(t, err)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.NUMBER, token.Type)
	assert.Equal(t, "123", token.Lexeme)
	assert.InDelta(t, 123.0, token.Literal, floatDelta)
	assert.Equal(t, 1, token.Line)
}

func TestScanToken_CommentThrownOut(t *testing.T) {
	var input string
	var scanner *Scanner
	var token *token.Token
	var err error

	input = "// hello!"
	scanner = NewScanner(input)
	verifyNextScanTokenIsEOF(t, scanner, 1)
	verifyNextScanTokenIsPastEOFError(t, scanner)

	input = "// hello!\n"
	scanner = NewScanner(input)
	verifyNextScanTokenIsEOF(t, scanner, 2)

	input = `some_ident
//a comment
while
`
	scanner = NewScanner(input)
	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "some_ident", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)

	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.WHILE, token.Type)
	assert.Equal(t, "while", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 3, token.Line)

	input = "abc//comment"
	scanner = NewScanner(input)
	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "abc", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)

	input = "a/true//comment"
	scanner = NewScanner(input)
	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.IDENTIFIER, token.Type)
	assert.Equal(t, "a", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.SLASH, token.Type)
	assert.Equal(t, "/", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	token, err = scanner.ScanToken()
	assert.Nil(t, err)
	assert.Equal(t, tokentype.TRUE, token.Type)
	assert.Equal(t, "true", token.Lexeme)
	assert.Equal(t, nil, token.Literal)
	assert.Equal(t, 1, token.Line)
	verifyNextScanTokenIsEOF(t, scanner, 1)

	input = "///comment"
	scanner = NewScanner(input)
	verifyNextScanTokenIsEOF(t, scanner, 1)

	input = `///comment1
//comment2
	//comment3`
	scanner = NewScanner(input)
	verifyNextScanTokenIsEOF(t, scanner, 3)

	input = `///comment1
//comment2
	//comment3
`
	scanner = NewScanner(input)
	verifyNextScanTokenIsEOF(t, scanner, 4)
}

func TestScanAllTokens_NoErrors(t *testing.T) {
	var tokens []*token.Token
	var errors []error

	input := ` hello
123
abc
 true false while
for    	while
	false


5*5
`
	scanner := NewScanner(input)
	tokens, errors = scanner.ScanAllTokens()
	assert.Len(t, tokens, 13) // The number of tokens in the string, plus an EOF token
	assert.Len(t, errors, 0)  // All valid input

	// Nothing returned when alling it again
	tokens, errors = scanner.ScanAllTokens()
	assert.Len(t, tokens, 0)
	assert.Len(t, errors, 0)
}

func TestScanAllTokens_RecoversFromErrors(t *testing.T) {
	var tokens []*token.Token
	var errors []error

	input := `abc 123 program
xyz ^^~[ _ {()()} ~


`
	scanner := NewScanner(input)
	tokens, errors = scanner.ScanAllTokens()
	assert.Len(t, tokens, 12) // The number of tokens in the string, plus an EOF token
	assert.Len(t, errors, 5)  // 6 unrecognized characters

	// Nothing returned when alling it again
	tokens, errors = scanner.ScanAllTokens()
	assert.Len(t, tokens, 0)
	assert.Len(t, errors, 0)
}

func TestScanAllTokens_IgnoresComments(t *testing.T) {
	input := `
	// here are some numbers
	1
	2
	// some more numbers
	3
	4
	// and finally, the last number
	5`
	scanner := NewScanner(input)
	tokens, errors := scanner.ScanAllTokens()
	assert.Len(t, tokens, 6) // The number of tokens in the string, plus an EOF token
	assert.Len(t, errors, 0) // All valid input
}
