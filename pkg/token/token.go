package token

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type Token struct {
	Type    tokentype.TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) String() string {
	if t.Literal == nil {
		return fmt.Sprintf("%s %s nil", t.Type, t.Lexeme)
	}
	return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
}
