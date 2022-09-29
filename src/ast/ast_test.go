package ast

import (
	"testing"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	p := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.New(token.LET, "let"),
				Name: &IdentifierExpression{
					Token: token.New(token.IDENT, "myVar"),
					Value: "myVar",
				},
				Value: &IdentifierExpression{
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
