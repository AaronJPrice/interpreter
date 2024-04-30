package parser

import (
	"errors"
	"testing"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/ast"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	t.Run("let x = 5;", func(t *testing.T) {
		input := `let x = 5;`

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.StatementLet{
					Token: token.New(token.LET, "let"),
					Name:  &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "x"), Value: "x"},
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
				&ast.StatementLet{
					Token: token.New(token.LET, "let"),
					Name:  &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "y"), Value: "y"},
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
				&ast.StatementLet{
					Token: token.New(token.LET, "let"),
					Name:  &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "foobar"), Value: "foobar"},
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
				&ast.StatementReturn{
					Token:       token.New(token.RETURN, "return"),
					ReturnValue: nil, // TODO
				},
				&ast.StatementReturn{
					Token:       token.New(token.RETURN, "return"),
					ReturnValue: nil, // TODO
				},
				&ast.StatementReturn{
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
			&ast.StatementExpression{
				Token: token.New(token.IDENT, "foobar"),
				Expression: &ast.ExpressionIdentifier{
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
			&ast.StatementExpression{
				Token: token.New(token.INT, "5"),
				Expression: &ast.ExpressionInteger{
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
				&ast.StatementExpression{
					Token: token.New(token.BANG, "!"),
					Expression: &ast.ExpressionPrefix{
						Token:    token.New(token.BANG, "!"),
						Operator: token.BANG,
						Right: &ast.ExpressionInteger{
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
				&ast.StatementExpression{
					Token: token.New(token.MINUS, "-"),
					Expression: &ast.ExpressionPrefix{
						Token:    token.New(token.MINUS, "-"),
						Operator: token.MINUS,
						Right: &ast.ExpressionInteger{
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
				&ast.StatementExpression{
					Token: token.New(token.INT, "5"),
					Expression: &ast.ExpressionInfix{
						Left: &ast.ExpressionInteger{
							Token: token.New(token.INT, "5"),
							Value: 5,
						},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right: &ast.ExpressionInteger{
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
				&ast.StatementExpression{
					Token: token.New(token.INT, "3"),
					Expression: &ast.ExpressionInfix{
						Left: &ast.ExpressionInteger{
							Token: token.New(token.INT, "3"),
							Value: 3,
						},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right: &ast.ExpressionInfix{
							Left: &ast.ExpressionInteger{
								Token: token.New(token.INT, "2"),
								Value: 2,
							},
							Token:    token.New(token.ASTERISK, "*"),
							Operator: token.ASTERISK,
							Right: &ast.ExpressionInteger{
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

func TestParseBoolean(t *testing.T) {
	t.Run("true;", func(t *testing.T) {
		input := "true;"

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.StatementExpression{
					Token: token.New(token.TRUE, "true"),
					Expression: &ast.ExpressionBoolean{
						Token: token.New(token.TRUE, "true"),
						Value: true,
					},
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("false;", func(t *testing.T) {
		input := "false;"

		expect := &ast.Program{
			Statements: []ast.Statement{
				&ast.StatementExpression{
					Token: token.New(token.FALSE, "false"),
					Expression: &ast.ExpressionBoolean{
						Token: token.New(token.FALSE, "false"),
						Value: false,
					},
				},
			},
		}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})
}

func TestParentheses(t *testing.T) {

	t.Run("1 + 2 + 3 + 4;", func(t *testing.T) {
		input := `1 + 2 + 3 + 4;`

		expect := &ast.Program{Statements: []ast.Statement{&ast.StatementExpression{
			Token: token.New(token.INT, "1"),
			Expression: &ast.ExpressionInfix{
				Left: &ast.ExpressionInfix{
					Left: &ast.ExpressionInfix{
						Left:     &ast.ExpressionInteger{Token: token.New(token.INT, "1"), Value: 1},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right:    &ast.ExpressionInteger{Token: token.New(token.INT, "2"), Value: 2},
					},
					Token:    token.New(token.PLUS, "+"),
					Operator: token.PLUS,
					Right:    &ast.ExpressionInteger{Token: token.New(token.INT, "3"), Value: 3},
				},
				Token:    token.New(token.PLUS, "+"),
				Operator: token.PLUS,
				Right:    &ast.ExpressionInteger{Token: token.New(token.INT, "4"), Value: 4},
			},
		}}}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})

	t.Run("1 + (2 + 3) + 4;", func(t *testing.T) {
		input := `1 + (2 + 3) + 4;`

		expect := &ast.Program{Statements: []ast.Statement{&ast.StatementExpression{
			Token: token.New(token.INT, "1"),
			Expression: &ast.ExpressionInfix{
				Left: &ast.ExpressionInfix{
					Left:     &ast.ExpressionInteger{Token: token.New(token.INT, "1"), Value: 1},
					Token:    token.New(token.PLUS, "+"),
					Operator: token.PLUS,
					Right: &ast.ExpressionInfix{
						Left:     &ast.ExpressionInteger{Token: token.New(token.INT, "2"), Value: 2},
						Token:    token.New(token.PLUS, "+"),
						Operator: token.PLUS,
						Right:    &ast.ExpressionInteger{Token: token.New(token.INT, "3"), Value: 3},
					},
				},
				Token:    token.New(token.PLUS, "+"),
				Operator: token.PLUS,
				Right:    &ast.ExpressionInteger{Token: token.New(token.INT, "4"), Value: 4},
			},
		}}}

		p := New(lexer.New(input))

		doTest(t, p, expect)
	})
}

func TestIf(t *testing.T) {
	t.Run(`if (x < y) { x; };`, func(t *testing.T) {
		input := `if (x < y) { x; };`

		expect := &ast.Program{Statements: []ast.Statement{&ast.StatementExpression{
			Token: token.New(token.IF, "if"),
			Expression: &ast.ExpressionIf{
				Token: token.New(token.IF, "if"),
				Condition: &ast.ExpressionInfix{
					Token:    token.New(token.LT, "<"),
					Left:     &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "x"), Value: "x"},
					Operator: token.LT,
					Right:    &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "y"), Value: "y"},
				},
				Consequence: &ast.StatementBlock{
					Token: token.New(token.LBRACE, "{"),
					Statements: []ast.Statement{&ast.StatementExpression{
						Token:      token.New(token.IDENT, "x"),
						Expression: &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "x"), Value: "x"},
					}},
				},
			},
		}}}

		p := New(lexer.New(input))
		doTest(t, p, expect)
	})

	t.Run(`if (x < y) { x; } else { y; };`, func(t *testing.T) {
		input := `if (x < y) { x; } else { y; };`

		expect := &ast.Program{Statements: []ast.Statement{&ast.StatementExpression{
			Token: token.New(token.IF, "if"),
			Expression: &ast.ExpressionIf{
				Token: token.New(token.IF, "if"),
				Condition: &ast.ExpressionInfix{
					Token:    token.New(token.LT, "<"),
					Left:     &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "x"), Value: "x"},
					Operator: token.LT,
					Right:    &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "y"), Value: "y"},
				},
				Consequence: &ast.StatementBlock{
					Token: token.New(token.LBRACE, "{"),
					Statements: []ast.Statement{&ast.StatementExpression{
						Token:      token.New(token.IDENT, "x"),
						Expression: &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "x"), Value: "x"},
					}},
				},
				Alternative: &ast.StatementBlock{
					Token: token.New(token.LBRACE, "{"),
					Statements: []ast.Statement{&ast.StatementExpression{
						Token:      token.New(token.IDENT, "y"),
						Expression: &ast.ExpressionIdentifier{Token: token.New(token.IDENT, "y"), Value: "y"},
					}},
				},
			},
		}}}

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
			t.Log("expect:", expect)
			t.Log("actual:", actual)
		}
	}
}
