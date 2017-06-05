// Package scanner implements the lexical analysis of S-expression grammar.
package scanner

import (
	//"fmt"

	"github.com/pravj/rugo/token"
	"github.com/pravj/rugo/utils"
)

// Scanner holds its states (metadata) during the analysis.
type Scanner struct {
	// input string to the scanner
	input string

	// slice of tokens collected by scanner
	Tokens []token.Token

	// starting position of in-process scan.
	// in-process scan means the hunt for the next lexeme
	// after finding a delimiter (whitespace).
	start int

	// current position of in-process scan
	current int

	// curent line number of in-process scan
	line int
}

// NewScanner returns a new scanner instance (*Scanner).
func NewScanner(input string) *Scanner {
	return &Scanner{
		input:  input,
		Tokens: make([]token.Token, 0),
		line:   1,
	}
}

// ScanTokens stops the scanning on finding the EOF token.
// It returns a slice of tokens representing the input string.
func (s *Scanner) ScanTokens() []token.Token {
	for !s.scanComplete() {
		s.start = s.current
		s.scanToken()
	}

	// append the terminal token (End Of File)
	s.Tokens = append(s.Tokens, token.New(token.EOF, "", s.line))

	return s.Tokens
}

// scanToken infers the lexical token from its string representation.
func (s *Scanner) scanToken() {
	nextChar := s.nextCharacter()

	switch nextChar {
	case "(":
		s.addToken(token.LEFT_PAREN)
		break
	case ")":
		s.addToken(token.RIGHT_PAREN)
		break
	case ":":
		for ((s.lookahead(0) != ",") && (s.lookahead(0) != ")")) && !s.scanComplete() {
			s.nextCharacter()
		}
		s.addToken(token.KEYWORD)
		break
	case ",":
		// ignore intentionally
		//s.addToken(token.COMMA)
		break
	case "+":
		s.addToken(token.PLUS)
		break
	case "-":
		s.addToken(token.MINUS)
		break
	case "*":
		s.addToken(token.STAR)
		break
	case "/":
		s.addToken(token.SLASH)
		break
	case "!":
		s.addConditionalToken("=", token.BANG_EQUAL, token.BANG)
		break
	case "=":
		s.addConditionalToken("=", token.EQUAL_EQUAL, token.EQUAL)
		break
	case ">":
		s.addConditionalToken("=", token.GREATER_EQUAL, token.GREATER)
		break
	case "<":
		s.addConditionalToken("=", token.LESS_EQUAL, token.LESS)
		break
	case " ":
		break
	case "\r":
		break
	case "\t":
		break
	case "\n":
		s.line++
		break
	case "\"":
		s.scanStringLiteral()
		break
	default:
		if utils.IsAlphaCharacter(nextChar) {
			s.scanNameLiteral()
		} else if utils.IsDigitCharacter(nextChar) {
			s.scanNumberLiteral()
		} else {
			s.addToken(token.INVALID)
		}
	}
}

// scanStringLiteral scans the prospective string-type token.
func (s *Scanner) scanStringLiteral() {
	for s.lookahead(0) != "\"" && !s.scanComplete() {
		if s.lookahead(0) == "\n" {
			s.line++
		}

		s.nextCharacter()
	}

	if s.scanComplete() {
		//errors.ReportError(s.line, s.lines[s.line-1], fmt.Sprintf("Unterminated string"))
		return
	}

	// the closing quote
	s.nextCharacter()

	// trim the surrounding quotes
	// TODO: deal with the literal value for a token (use interface)
	// fmt.Println(string(s.input[s.start+1:s.current-1]))
	s.addToken(token.STRING)
}

// scanNameLiteral scans the prospective name-type token.
func (s *Scanner) scanNameLiteral() {
	for utils.IsAlphaCharacter(s.lookahead(0)) && !s.scanComplete() {
		s.nextCharacter()
	}

	s.addToken(token.NAME)
}

// scanNumberPart scans the sections of a number-type token.
func (s *Scanner) scanNumberPart() {
	for utils.IsDigitCharacter(s.lookahead(0)) && !s.scanComplete() {
		s.nextCharacter()
	}
}

// scanNumberLiteral scans the prospective number-type token.
func (s *Scanner) scanNumberLiteral() {
	// scan the numbers (0-9)+
	s.scanNumberPart()

	// if the number has a decimal part also
	if s.matchCharacter(".") {
		if utils.IsDigitCharacter(s.lookahead(0)) {
			s.scanNumberPart()
		} else {
			// a number with a 'decimal' without prefix part is 'invalid' ^[0-9]+\.$
			s.addToken(token.INVALID)
			return
		}
	}

	s.addToken(token.NUMBER)
}

// addToken appends the given token to the scanner's tokens.
func (s *Scanner) addToken(tokenType token.TokenType) {
	lexemeStringValue := string(s.input[s.start:s.current])
	s.Tokens = append(s.Tokens, token.New(tokenType, lexemeStringValue, s.line))
}

// addConditionalToken appends one of the given tokens to scanner's tokens.
func (s *Scanner) addConditionalToken(expectedChar string, trueToken, falseToken token.TokenType) {
	// if the next character is as expected
	if s.matchCharacter(expectedChar) {
		s.addToken(trueToken) // trueToken
	} else {
		s.addToken(falseToken) // falseToken
	}
}

// nextCharacter returns the current character from input string.
func (s *Scanner) nextCharacter() string {
	s.current++
	return string(s.input[s.current-1])
}

// matchCharacter checks if the next character is as expected.
// It consumes the next character only if it matches as expected.
func (s *Scanner) matchCharacter(expectedChar string) bool {
	if s.scanComplete() {
		return false
	}

	if string(s.input[s.current]) != expectedChar {
		return false
	}

	s.current++
	return true
}

// lookahead returns the next character on a given position
// from the current location of scanner.
func (s *Scanner) lookahead(length int) string {
	if s.current+length >= len(s.input) {
		return "\000"
	}

	return string(s.input[s.current+length])
}

// scanComplete checks if the scanner has finished entire input string.
// It returns true if successful.
func (s *Scanner) scanComplete() bool {
	return s.current >= len(s.input)
}
