package ast

import (
	"bytes"
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"
)

type StatementLet struct {
	Token token.Token // the token.LET token
	Name  *ExpressionIdentifier
	Value Expression
}

func (s *StatementLet) statementNode()       {}
func (s *StatementLet) TokenLiteral() string { return s.Token.Literal }
func (s *StatementLet) String() string {
	return fmt.Sprintf("%v %v = %v;", s.TokenLiteral(), s.Name.String(), s.Value.String())
}

type StatementReturn struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (s *StatementReturn) statementNode()       {}
func (s *StatementReturn) TokenLiteral() string { return s.Token.Literal }
func (s *StatementReturn) String() string {
	return fmt.Sprintf("%v %v", s.TokenLiteral(), s.ReturnValue)
}

type StatementExpression struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (s *StatementExpression) statementNode()       {}
func (s *StatementExpression) TokenLiteral() string { return s.Token.Literal }
func (s *StatementExpression) String() string {
	return s.Expression.String()
}

type StatementBlock struct {
	Token      token.Token
	Statements []Statement
}

func (s *StatementBlock) statementNode()       {}
func (s *StatementBlock) TokenLiteral() string { return s.Token.Literal }
func (s *StatementBlock) String() string {
	var out bytes.Buffer
	for _, s := range s.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
