package parser

import (
	"errors"
	"testing"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	t.Run("let x = 5;", func(t *testing.T) {
		input := `let x = 5;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.New(token.LET, "let"),
					Name:  &ast.IdentifierExpression{Token: token.New(token.IDENT, "x"), Value: "x"},
					Value: nil, // TODO
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("let y = 10;", func(t *testing.T) {
		input := `let y = 10;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.New(token.LET, "let"),
					Name:  &ast.IdentifierExpression{Token: token.New(token.IDENT, "y"), Value: "y"},
					Value: nil, // TODO
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("let foobar = 838383;", func(t *testing.T) {
		input := `let foobar = 838383;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.New(token.LET, "let"),
					Name:  &ast.IdentifierExpression{Token: token.New(token.IDENT, "foobar"), Value: "foobar"},
					Value: nil, // TODO
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("let x 5;", func(t *testing.T) {
		input := `let x 5;`

		expect := []error{
			errors.New("expected next token to be =, got INT"),
		}

		p := New(lexer.New(input))
		p.ParseProgram()
		actual := p.Errors()

		assert.Equal(t, expect, actual)
	})
}

func TestReturnStatements(t *testing.T) {
	t.Run("return 9; return 10; return 12094;", func(t *testing.T) {
		input := `
		return 9;
		return 10;
		return 12094;
		`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Token:       token.New(token.RETURN, "return"),
					ReturnValue: nil, // TODO
				},
				&ast.ReturnStatement{
					Token:       token.New(token.RETURN, "return"),
					ReturnValue: nil, // TODO
				},
				&ast.ReturnStatement{
					Token:       token.New(token.RETURN, "return"),
					ReturnValue: nil, // TODO
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})
}

func doTest(t *testing.T, p *Parser, expect interface{}) {
	actual := p.ParseProgram()

	if errs := p.Errors(); !assert.Nil(t, errs) {
		for _, err := range errs {
			t.Log(err)
		}
	} else {
		assert.Equal(t, expect, actual)
	}
}
