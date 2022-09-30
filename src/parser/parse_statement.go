package parser

import (
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.crntToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}
func (p *Parser) parseLetStatement() *ast.LetStatement {
	s := &ast.LetStatement{Token: p.crntToken}

	if !p.advanceIfNextIs(token.IDENT) {
		return nil // could return an error here
	}

	s.Name = &ast.IdentifierExpression{Token: p.crntToken, Value: p.crntToken.Literal}

	if !p.advanceIfNextIs(token.ASSIGN) {
		return nil // could return an error here
	}

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	s := &ast.ReturnStatement{Token: p.crntToken}

	p.advance()

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	s := &ast.ExpressionStatement{
		Token:      p.crntToken,
		Expression: p.parseExpression(LOWEST),
	}

	p.advanceIfNextIs(token.SEMICOLON)

	return s
}
