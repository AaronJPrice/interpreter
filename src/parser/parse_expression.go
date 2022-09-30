package parser

import (
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {

	switch p.crntToken.Type {
	case token.IDENT:
		return p.parseIdentifier()
	default:
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
