package loxerr

import (
	"fmt"

	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
)

func getWhere(t *token.Token) string {
	if t.Type == tokentype.EOF {
		return "at end"
	}
	return fmt.Sprintf("at '%s'", t.Lexeme)
}

type LoxErrorAtToken struct {
	Token   *token.Token
	where   string
	message string
}

func AtToken(t *token.Token, message string) *LoxErrorAtToken {
	return &LoxErrorAtToken{
		Token:   t,
		where:   getWhere(t),
		message: message,
	}
}

func (e *LoxErrorAtToken) Error() string {
	return fmt.Sprintf("[line %d] Error %s: %s", e.Token.Line, e.where, e.message)
}

type LoxErrorAtLine struct {
	line    int
	message string
}

func AtLine(line int, message string) *LoxErrorAtLine {
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

func Internal(message string) *LoxInternalError {
	return &LoxInternalError{message}
}

func (e *LoxInternalError) Error() string {
	return fmt.Sprintf("Internal error: %s", e.message)
}

type LoxRuntimeError struct {
	Token   *token.Token
	where   string
	message string
}

func Runtime(t *token.Token, message string) *LoxRuntimeError {
	return &LoxRuntimeError{
		Token:   t,
		where:   getWhere(t),
		message: message,
	}
}

func (e *LoxRuntimeError) Error() string {
	return fmt.Sprintf("[line %d] Runtime error %s: %s", e.Token.Line, e.where, e.message)
}
