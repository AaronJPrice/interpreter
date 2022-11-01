package ast

import (
	"bytes"
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

type ExpressionIdentifier struct {
	Token token.Token
	Value string
}

func (e *ExpressionIdentifier) expressionNode()      {}
func (e *ExpressionIdentifier) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionIdentifier) String() string       { return e.Value }

type ExpressionInteger struct {
	Token token.Token
	Value int64
}

func (il *ExpressionInteger) expressionNode()      {}
func (il *ExpressionInteger) TokenLiteral() string { return il.Token.Literal }
func (il *ExpressionInteger) String() string       { return il.Token.Literal }

type ExpressionPrefix struct {
	Token    token.Token
	Operator token.TokenType
	Right    Expression
}

func (e *ExpressionPrefix) expressionNode()      {}
func (e *ExpressionPrefix) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionPrefix) String() string {
	return fmt.Sprintf("(%v%v)", e.Operator, e.Right.String())
}

type ExpressionInfix struct {
	Token    token.Token
	Left     Expression
	Operator token.TokenType
	Right    Expression
}

func (e *ExpressionInfix) expressionNode()      {}
func (e *ExpressionInfix) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionInfix) String() string {
	return fmt.Sprintf("(%v %v %v)", e.Left.String(), e.Operator, e.Right.String())
}

type ExpressionBoolean struct {
	Token token.Token
	Value bool
}

func (e *ExpressionBoolean) expressionNode()      {}
func (e *ExpressionBoolean) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionBoolean) String() string       { return e.Token.Literal }

type ExpressionIf struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *StatementBlock
	Alternative *StatementBlock
}

func (e *ExpressionIf) expressionNode()      {}
func (e *ExpressionIf) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionIf) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())
	if e.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(e.Alternative.String())
	}
	return out.String()
}
