package ast

import (
	"testing"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	p := &Program{
		Statements: []Statement{
			&StatementLet{
				Token: token.New(token.LET, "let"),
				Name: &ExpressionIdentifier{
					Token: token.New(token.IDENT, "myVar"),
					Value: "myVar",
				},
				Value: &ExpressionIdentifier{
					Token: token.New(token.IDENT, "anotherVar"),
					Value: "anotherVar",
				},
			},
		},
	}
	expect := "let myVar = anotherVar;"
	actual := p.String()
	assert.Equal(t, expect, actual)
}
