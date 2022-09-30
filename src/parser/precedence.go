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

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) nextPrecedence() int {
	return p.getPrecedence(p.nextToken)
}

func (p *Parser) crntPrecedence() int {
	return p.getPrecedence(p.crntToken)
}

func (p *Parser) getPrecedence(t token.Token) int {
	if p, ok := precedences[t.Type]; ok {
		return p
	}
	return LOWEST
}
