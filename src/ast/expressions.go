package ast

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

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

type PrefixExpression struct {
	Token    token.Token
	Operator token.TokenType
	Right    Expression
}

func (e *PrefixExpression) expressionNode()      {}
func (e *PrefixExpression) TokenLiteral() string { return e.Token.Literal }
func (e *PrefixExpression) String() string {
	return fmt.Sprintf("(%v%v)", e.Operator, e.Right.String())
}
