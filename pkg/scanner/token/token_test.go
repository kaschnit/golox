package token

import (
	"testing"

	"github.com/kaschnit/golox/pkg/scanner/token/tokentype"
	"github.com/stretchr/testify/assert"
)

func TestTokenToString(t *testing.T) {
	var token Token

	token = Token{
		tokentype.ELSE,
		"abc",
		nil,
		3,
	}
	assert.Equal(t, "ELSE abc nil", token.String())

	token = Token{
		tokentype.FOR,
		"for",
		55,
		3,
	}
	assert.Equal(t, "FOR for 55", token.String())

	token = Token{
		tokentype.BANG_EQUAL,
		"123",
		"xyz",
		3,
	}
	assert.Equal(t, "BANG_EQUAL 123 xyz", token.String())

	token = Token{
		tokentype.EOF,
		"",
		nil,
		3,
	}
	assert.Equal(t, "EOF  nil", token.String())
}
