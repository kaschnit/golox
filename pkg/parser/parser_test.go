package parser

import (
	"testing"

	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

func strToken(val string) *token.Token {
	return &token.Token{
		Type:    tokentype.STRING,
		Lexeme:  val,
		Literal: val,
	}
}

func TestParseExpression_Equality(t *testing.T) {
	// parser := NewParser([]*token.Token{
	// 	makeToken(tokentype.NUMBER)
	// })
}
