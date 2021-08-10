package loxerr

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

type LoxErrorAtToken struct {
	token   *token.Token
	where   string
	message string
}

func NewLoxErrorAtToken(t *token.Token, message string) *LoxErrorAtToken {
	var where string
	if t.Type == tokentype.EOF {
		where = " at end"
	} else {
		where = fmt.Sprintf(" at '%s'", t.Lexeme)
	}

	return &LoxErrorAtToken{
		token:   t,
		where:   where,
		message: message,
	}
}

func (e *LoxErrorAtToken) Error() string {
	return fmt.Sprintf("[line %d] Error%s: %s", e.token.Line, e.where, e.message)
}

type LoxErrorAtLine struct {
	line    int
	message string
}

func NewLoxErrorAtLine(line int, message string) *LoxErrorAtLine {
	return &LoxErrorAtLine{
		line:    line,
		message: message,
	}
}

func (e *LoxErrorAtLine) Error() string {
	return fmt.Sprintf("[line %d] Error: %s", e.line, e.message)
}

type LoxInternalError struct {
	message string
}

func NewLoxInternalError(message string) *LoxInternalError {
	return &LoxInternalError{message}
}

func (e *LoxInternalError) Error() string {
	return fmt.Sprintf("Internal error: %s", e.message)
}
