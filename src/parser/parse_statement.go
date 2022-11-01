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

func (p *Parser) parseLetStatement() *ast.StatementLet {
	s := &ast.StatementLet{Token: p.crntToken}

	if !p.advanceIfNextIs(token.IDENT) {
		return nil // could return an error here
	}

	s.Name = &ast.ExpressionIdentifier{Token: p.crntToken, Value: p.crntToken.Literal}

	if !p.advanceIfNextIs(token.ASSIGN) {
		return nil // could return an error here
	}

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
}

func (p *Parser) parseReturnStatement() *ast.StatementReturn {
	s := &ast.StatementReturn{Token: p.crntToken}

	p.advance()

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
}

func (p *Parser) parseExpressionStatement() *ast.StatementExpression {
	s := &ast.StatementExpression{Token: p.crntToken}
	// parseExpression advances crntToken, so it must be called _after_ setting the Token field in
	// ExpressionStatement (i.e. it can't be called inline as part of the struct literal) since
	// otherwise the value of Token will be wrong
	s.Expression = p.parseExpression(LOWEST)

	p.advanceIfNextIs(token.SEMICOLON)

	return s
}
