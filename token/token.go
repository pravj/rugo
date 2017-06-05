// Package token defines the lexical tokens (lexems) for S-expression grammar.
package token

// TokenType represents a set of lexical tokens.
type TokenType int

const (
	EOF TokenType = iota
	INVALID

	LEFT_PAREN
	RIGHT_PAREN

	COMMA
	//COLON

	STRING  // alphabates with a quote around
	NAME    // alphabates without a quote around
	KEYWORD // NAME WITH A COLON (':') PREFIX
	NUMBER  // [0-9]+(\.[0-9]+)?

	PLUS
	MINUS
	STAR
	SLASH

	BANG
	EQUAL
	GREATER
	LESS

	BANG_EQUAL
	EQUAL_EQUAL
	GREATER_EQUAL
	LESS_EQUAL
)

// Token represents metadata about a lexical token
type Token struct {
	// type of the lexical token
	TypeOfToken TokenType

	// string represents of the lexical token
	Lexeme string

	// line number for the token
	line int
}

// New returns a new Token
func New(typeOfToken TokenType, lexeme string, line int) Token {
	return Token{TypeOfToken: typeOfToken, Lexeme: lexeme, line: line}
}
