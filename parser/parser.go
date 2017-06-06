/*
Package parser implements the top-down parsing of tokenised input string.
*/
package parser

import (
	"github.com/pravj/rugo/ast"
	"github.com/pravj/rugo/token"
)

// Parser represents the internal state of the parser
type Parser struct {
	// tokens generated from the scanner
	tokens []token.Token

	// index of the next token to parse
	current int
}

// NewParser returns a new instance of Parser (*Parser)
func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// parse starts the process of top-down parsing.
func (p *Parser) Parse() ast.Node {
	t := p.advanceToken()

	if (t.TypeOfToken == token.S_EXP_OPEN) {
		// initialize a new node list
		NodeList := ast.NewNode()
		if NodeList.Name == "" {
			NodeList.Name = p.peek(0).Lexeme
		}

		// don't add self to node children
		p.advanceToken()

		for !p.checkToken(0, token.RIGHT_PAREN) {
			NodeList.Nodes = append(NodeList.Nodes, p.Parse())
		}

		// parses the closing bracet ')'
		p.advanceToken()

		return NodeList
	} else if t.Lexeme == ")" {
		panic("Unexpected closing bracet ')'.")
	} else {
		node := ast.Node{Name: t.Lexeme}

		return node
	}
}

// advanceToken consumes the current token and returns it.
func (p *Parser) advanceToken() token.Token {
	if !p.parseComplete() {
		p.current++
	}

	return p.previous()
}

// previous returns the recently consumed token.
func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

// parseComplete checks if we have run out of tokens.
func (p *Parser) parseComplete() bool {
	return p.peek(0).TypeOfToken == token.EOF
}

// peek returns the current token we have yet to consume.
func (p *Parser) peek(length int) token.Token {
	return p.tokens[p.current+length]
}

// checkToken checks if token at given offset from current is of a given type.
func (p *Parser) checkToken(length int, token token.TokenType) bool {
	if p.current+length >= len(p.tokens) {
		return false
	}

	return p.peek(length).TypeOfToken == token
}
