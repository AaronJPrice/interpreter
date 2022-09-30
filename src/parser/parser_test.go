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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	expect := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token: token.New(token.IDENT, "foobar"),
				Expression: &ast.IdentifierExpression{
					Token: token.New(token.IDENT, "foobar"),
					Value: "foobar",
				},
			},
		},
	}

	p := New(lexer.New(input))

	doTest(t, p, expect)
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	expect := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token: token.New(token.INT, "5"),
				Expression: &ast.IntegerExpression{
					Token: token.New(token.INT, "5"),
					Value: 5,
				},
			},
		},
	}

	p := New(lexer.New(input))

	doTest(t, p, expect)
}

func TestParsingPrefixExpressions(t *testing.T) {
	t.Run("!5;", func(t *testing.T) {

		input := `!5;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Token: token.New(token.BANG, "!"),
					Expression: &ast.PrefixExpression{
						Token:    token.New(token.BANG, "!"),
						Operator: token.BANG,
						Right: &ast.IntegerExpression{
							Token: token.New(token.INT, "5"),
							Value: 5,
						},
					},
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("-15;", func(t *testing.T) {

		input := `-15;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Token: token.New(token.MINUS, "-"),
					Expression: &ast.PrefixExpression{
						Token:    token.New(token.MINUS, "-"),
						Operator: token.MINUS,
						Right: &ast.IntegerExpression{
							Token: token.New(token.INT, "15"),
							Value: 15,
						},
					},
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})
}

func TestParseInfixOperators(t *testing.T) {
	t.Run("5 + 5;", func(t *testing.T) {
		input := `5 + 5;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Token: token.New(token.INT, "5"),
					Expression: &ast.InfixExpression{
						Left: &ast.IntegerExpression{
							Token: token.New(token.INT, "5"),
							Value: 5,
						},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right: &ast.IntegerExpression{
							Token: token.New(token.INT, "5"),
							Value: 5,
						},
					},
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("3 + 2 * 1;", func(t *testing.T) {
		input := `3 + 2 * 1;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Token: token.New(token.INT, "3"),
					Expression: &ast.InfixExpression{
						Left: &ast.IntegerExpression{
							Token: token.New(token.INT, "3"),
							Value: 3,
						},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right: &ast.InfixExpression{
							Left: &ast.IntegerExpression{
								Token: token.New(token.INT, "2"),
								Value: 2,
							},
							Token:    token.New(token.ASTERISK, "*"),
							Operator: token.ASTERISK,
							Right: &ast.IntegerExpression{
								Token: token.New(token.INT, "1"),
								Value: 1,
							},
						},
					},
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
		if !assert.Equal(t, expect, actual) {
			t.Log("actual:", actual)
		}
	}
}
