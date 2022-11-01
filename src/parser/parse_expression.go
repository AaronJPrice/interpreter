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
	case token.TRUE, token.FALSE:
		expression = p.parseBoolean()
	case token.BANG, token.MINUS:
		expression = p.parsePrefixExpression()
	case token.LPAREN:
		expression = p.parseGroupedExpression()
	case token.IF:
		expression = p.parseIfExpression()
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
	return &ast.ExpressionIdentifier{Token: p.crntToken, Value: p.crntToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.crntToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("could not parse %q as integer", p.crntToken.Literal))
		return nil
	}

	return &ast.ExpressionInteger{
		Token: p.crntToken,
		Value: value,
	}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.ExpressionBoolean{
		Token: p.crntToken,
		Value: p.crntToken.Type == token.TRUE,
	}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.ExpressionPrefix{
		Token:    p.crntToken,
		Operator: p.crntToken.Type,
	}
	p.advance()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.ExpressionInfix{
		Left:     left,
		Token:    p.crntToken,
		Operator: p.crntToken.Type,
	}
	precedence := p.crntPrecedence()
	p.advance()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()
	expression := p.parseExpression(LOWEST)
	if !p.advanceIfNextIs(token.RPAREN) {
		return nil
	}
	return expression
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.ExpressionIf{Token: p.crntToken}

	if !p.advanceIfNextIs(token.LPAREN) {
		return nil
	}
	p.advance()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.advanceIfNextIs(token.RPAREN) {
		return nil
	}
	if !p.advanceIfNextIs(token.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.nextToken.Type == token.ELSE {
		p.advance()
		if !p.advanceIfNextIs(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Errorf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
