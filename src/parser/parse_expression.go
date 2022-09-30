package parser

import (
	"fmt"
	"strconv"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {

	var expression ast.Expression

	switch p.crntToken.Type {
	case token.IDENT:
		expression = p.parseIdentifier()
	case token.INT:
		expression = p.parseIntegerLiteral()
	case token.BANG, token.MINUS:
		expression = p.parsePrefixExpression()
	default:
		p.noPrefixParseFnError(p.crntToken.Type)
		return nil
	}

	for p.nextToken.Type != token.SEMICOLON && precedence < p.nextPrecedence() {
		switch p.nextToken.Type {
		case token.PLUS, token.MINUS, token.SLASH, token.ASTERISK, token.EQ, token.NOT_EQ, token.LT, token.GT:
			p.advance()
			expression = p.parseInfixExpression(expression)
		default:
			return expression
		}
	}

	return expression
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

	return &ast.IntegerExpression{
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

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Left:     left,
		Token:    p.crntToken,
		Operator: p.crntToken.Type,
	}
	precedence := p.crntPrecedence()
	p.advance()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Errorf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
