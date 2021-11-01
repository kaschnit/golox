package scanner

import (
	"fmt"
	"strconv"

	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
	"github.com/kaschnit/golox/pkg/utils/stringutil"
)

type Scanner struct {
	// The source code to tokenize.
	source []rune

	// Whether or not any errors have been encountered while scanning so far.
	hasError bool

	// Tracks whether the entire input, including EOF, has already been scanned.
	// This prevents scannign the same EOF at the end of the input repeatedly.
	finished bool

	// Points to the start of the current lexeme currently being tokenized.
	start int

	// Points to the current character of the lexeme currently being tokenized.
	current int

	// Line number of the lexeme being tokenized.
	line int
}

// Create a Scanner instance.
func NewScanner(source string) *Scanner {
	return &Scanner{
		source:   []rune(source),
		hasError: false,
		finished: false,
		start:    0,
		current:  0,
		line:     1,
	}
}

// Reset the scanner to initial state.
func (s *Scanner) Reset() {
	s.hasError = false
	s.finished = false
	s.start = 0
	s.current = 0
	s.line = 1
}

// Tokenize the remaining input that has not been scanned yet.
func (s *Scanner) ScanAllTokens() ([]*token.Token, error) {
	tokens := make([]*token.Token, 0)
	errs := make([]error, 0)
	for !s.finished {
		nextToken, err := s.ScanToken()

		// Continue scanning even if an error is encountered.
		if err != nil {
			errs = append(errs, err)
		} else {
			tokens = append(tokens, nextToken)
		}

	}

	if len(errs) > 0 {
		return tokens, loxerr.NewLoxMultiError(errs)
	}

	return tokens, nil
}

// Scan the next token.
func (s *Scanner) ScanToken() (*token.Token, error) {
	if s.finished {
		return nil, loxerr.NewLoxInternalError("Scanner has already reached EOF")
	}

	s.start = s.current
	result, err := s.scanToken()

	// Keep moving on if there's no errors but no token is returned.
	// Do this to handle whitespace that won't matter after tokenization.
	for err == nil && result == nil {
		s.start = s.current
		result, err = s.scanToken()
	}

	return result, err
}

// Helper for scanning one token.
// Returns nil for token if a whitespace is found.
func (s *Scanner) scanToken() (*token.Token, error) {
	if s.isAtEnd() {
		s.finished = true
		return &token.Token{
			Type:    tokentype.EOF,
			Lexeme:  "",
			Literal: nil,
			Line:    s.line,
		}, nil
	}

	char := s.advance()

	switch char {
	// Tokens guaranteed to be 1 character
	case '(':
		return s.createToken(tokentype.LEFT_PAREN), nil
	case ')':
		return s.createToken(tokentype.RIGHT_PAREN), nil
	case '{':
		return s.createToken(tokentype.LEFT_BRACE), nil
	case '}':
		return s.createToken(tokentype.RIGHT_BRACE), nil
	case ',':
		return s.createToken(tokentype.COMMA), nil
	case '.':
		return s.createToken(tokentype.DOT), nil
	case '-':
		return s.createToken(tokentype.MINUS), nil
	case '+':
		return s.createToken(tokentype.PLUS), nil
	case ';':
		return s.createToken(tokentype.SEMICOLON), nil
	case '*':
		return s.createToken(tokentype.STAR), nil

	// Could be 1 or 2 character tokens
	case '/':
		if s.peek(1) == '/' {
			s.throwAwayLine()
			return nil, nil
		} else {
			return s.createToken(tokentype.SLASH), nil
		}
	case '!':
		if s.peek(1) == '=' {
			s.current++
			return s.createToken(tokentype.BANG_EQUAL), nil
		} else {
			return s.createToken(tokentype.BANG), nil
		}
	case '>':
		if s.peek(1) == '=' {
			s.current++
			return s.createToken(tokentype.GREATER_EQUAL), nil
		} else {
			return s.createToken(tokentype.GREATER), nil
		}
	case '<':
		if s.peek(1) == '=' {
			s.current++
			return s.createToken(tokentype.LESS_EQUAL), nil
		} else {
			return s.createToken(tokentype.LESS), nil
		}
	case '=':
		if s.peek(1) == '=' {
			s.current++
			return s.createToken(tokentype.EQUAL_EQUAL), nil
		} else {
			return s.createToken(tokentype.EQUAL), nil
		}

	// Ignore whitespace
	case '\n':
		s.line++
		fallthrough
	case ' ':
		fallthrough
	case '\r':
		fallthrough
	case '\t':
		s.start = s.current
		return nil, nil

	// Strings
	case '"':
		return s.scanString()

	// All other tokens that are of arbitrary length
	default:
		if stringutil.IsRuneNumeric(char) {
			return s.scanNumber()
		} else if stringutil.IsRuneAlpha(char) {
			return s.scanIdentifier()
		}
		s.hasError = true
		errMsg := fmt.Sprintf("Unrecognized character %s", string(char))
		return nil, loxerr.NewLoxErrorAtLine(s.line, errMsg)
	}
}

func (s *Scanner) scanString() (*token.Token, error) {
	// Advance the current pointer until EOF or closing quote is encountered.
	for s.peek(1) != '"' && !s.isAtEnd() {
		if s.peek(1) == '\n' {
			s.line++
		}
		s.current++
	}

	if s.isAtEnd() {
		return nil, loxerr.NewLoxErrorAtLine(s.line, "Unterminated string.")
	}

	// Increment the current pointer past the closing '"'
	s.current++

	return &token.Token{
		Type:    tokentype.STRING,
		Lexeme:  s.currentLexeme(),
		Literal: s.subCurrentLexeme(1, 1),
		Line:    s.line,
	}, nil
}

func (s *Scanner) scanNumber() (*token.Token, error) {
	for stringutil.IsRuneNumeric(s.peek(1)) {
		s.current++
	}
	if s.peek(1) == '.' && stringutil.IsRuneNumeric(s.peek(2)) {
		s.current++
		for stringutil.IsRuneNumeric(s.peek(1)) {
			s.current++
		}
	}

	lexeme := s.currentLexeme()
	literal, err := strconv.ParseFloat(lexeme, 32)
	if err != nil {
		return nil, err
	}

	return &token.Token{
		Type:    tokentype.NUMBER,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    s.line,
	}, nil
}

func (s *Scanner) scanIdentifier() (*token.Token, error) {
	for stringutil.IsRuneAlphaNumeric(s.peek(1)) {
		s.current++
	}

	lexeme := s.currentLexeme()
	return &token.Token{
		Type:    tokentype.FromIdentifier(lexeme),
		Lexeme:  lexeme,
		Literal: nil,
		Line:    s.line,
	}, nil
}

func (s *Scanner) throwAwayLine() {
	// Advance to the next newline or until EOF.
	for s.peek(1) != '\n' && !s.isAtEnd() {
		s.current++
	}
	s.start = s.current
}

// Helper for creating a token based on the scanner's current state.
func (s *Scanner) createToken(tokenType tokentype.TokenType) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Lexeme:  s.currentLexeme(),
		Literal: nil,
		Line:    s.line,
	}
}

// Get the lexeme that the scanner is currently pointing to.
func (s *Scanner) currentLexeme() string {
	return string(s.source[s.start:s.current])
}

// Get a substring of the lexeme that the scanner is currently pointing to.
func (s *Scanner) subCurrentLexeme(beginOffset int, endOffset int) string {
	return string(s.source[s.start+beginOffset : s.current-endOffset])
}

// Get the character that is lookahead in front of the current pointer.
func (s *Scanner) peek(lookahead int) rune {
	// Return null char if the scanner is at the end of input
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current+lookahead-1]
}

// Advance the current pointer to the next lexeme.
func (s *Scanner) advance() rune {
	result := s.peek(1)
	s.current++
	return result
}

// Whether or not the scanner has tokenized all input.
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
