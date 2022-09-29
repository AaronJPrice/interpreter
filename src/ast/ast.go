package ast

import (
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

////////////////////////////////////

type IdentifierExpression struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (e *IdentifierExpression) expressionNode()      {}
func (e *IdentifierExpression) TokenLiteral() string { return e.Token.Literal }
func (e *IdentifierExpression) String() string {
	return e.Value
}
