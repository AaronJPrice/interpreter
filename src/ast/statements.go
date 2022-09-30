package ast

import (
	"bytes"
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *IdentifierExpression
	Value Expression
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LetStatement) String() string {
	return fmt.Sprintf("%v %v = %v;", s.TokenLiteral(), s.Name.String(), s.Value.String())
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (s *ReturnStatement) statementNode()       {}
func (s *ReturnStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ReturnStatement) String() string {
	return fmt.Sprintf("%v %v", s.TokenLiteral(), s.ReturnValue)
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (s *ExpressionStatement) statementNode()       {}
func (s *ExpressionStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ExpressionStatement) String() string {
	return s.Expression.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (s *BlockStatement) statementNode()       {}
func (s *BlockStatement) TokenLiteral() string { return s.Token.Literal }
func (s *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range s.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
