package parser

import "bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"

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
	return p.getPrecedence(p.nextToken)
}

func (p *Parser) crntPrecedence() int {
	return p.getPrecedence(p.crntToken)
}

func (p *Parser) getPrecedence(t token.Token) int {
	switch t.Type {
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
