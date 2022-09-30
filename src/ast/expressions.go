package ast

import (
	"bytes"
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

type IntegerExpression struct {
	Token token.Token
	Value int64
}

func (il *IntegerExpression) expressionNode()      {}
func (il *IntegerExpression) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerExpression) String() string       { return il.Token.Literal }

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

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator token.TokenType
	Right    Expression
}

func (e *InfixExpression) expressionNode()      {}
func (e *InfixExpression) TokenLiteral() string { return e.Token.Literal }
func (e *InfixExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.Left.String(), e.Operator, e.Right.String())
}

type BooleanExpression struct {
	Token token.Token
	Value bool
}

func (e *BooleanExpression) expressionNode()      {}
func (e *BooleanExpression) TokenLiteral() string { return e.Token.Literal }
func (e *BooleanExpression) String() string       { return e.Token.Literal }

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
