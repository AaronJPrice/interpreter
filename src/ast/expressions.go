package ast

import "bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"

type IdentifierExpression struct {
	Token token.Token
	Value string
}

func (e *IdentifierExpression) expressionNode()      {}
func (e *IdentifierExpression) TokenLiteral() string { return e.Token.Literal }
func (e *IdentifierExpression) String() string       { return e.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
