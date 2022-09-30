package parser

import (
	"fmt"
	"strconv"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {

	switch p.crntToken.Type {
	case token.IDENT:
		return p.parseIdentifier()
	case token.INT:
		return p.parseIntegerLiteral()
	case token.BANG, token.MINUS:
		return p.parsePrefixExpression()
	default:
		p.noPrefixParseFnError(p.crntToken.Type)
		return nil
	}

	// if prefixParseFn, exists := p.prefixParseFns[p.crntToken.Type]; exists {
	// 	return prefixParseFn()
	// }

	// return nil
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.IdentifierExpression{Token: p.crntToken, Value: p.crntToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.crntToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("could not parse %q as integer", p.crntToken.Literal))
		return nil
	}

	return &ast.IntegerLiteral{
		Token: p.crntToken,
		Value: value,
	}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.crntToken,
		Operator: p.crntToken.Type,
	}
	p.advance()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Errorf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
