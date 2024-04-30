package parser

import "bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (p *Parser) nextPrecedence() int {
	return p.getPrecedence(p.nextToken.Type)
}

func (p *Parser) crntPrecedence() int {
	return p.getPrecedence(p.crntToken.Type)
}

func (p *Parser) getPrecedence(t token.TokenType) int {
	switch t {
	case token.EQ, token.NOT_EQ:
		return EQUALS
	case token.LT, token.GT:
		return LESSGREATER
	case token.PLUS, token.MINUS:
		return SUM
	case token.SLASH, token.ASTERISK:
		return PRODUCT
	default:
		return LOWEST
	}
}
