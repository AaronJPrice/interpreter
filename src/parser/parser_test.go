package parser

import (
	"errors"
	"testing"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {

	t.Run("let statements", func(t *testing.T) {
		t.Run("let x = 5;", func(t *testing.T) {
			input := `let x = 5;`

			expect := &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.New(token.LET, "let"),
						Name:  &ast.Identifier{Token: token.New(token.IDENT, "x"), Value: "x"},
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
						Name:  &ast.Identifier{Token: token.New(token.IDENT, "y"), Value: "y"},
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
						Name:  &ast.Identifier{Token: token.New(token.IDENT, "foobar"), Value: "foobar"},
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
	})

}

// func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
// 	if s.TokenLiteral() != "let" {
// 		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
// 		return false
// 	}
// 	letStmt, ok := s.(*ast.LetStatement)
// 	if !ok {
// 		t.Errorf("s not *ast.LetStatement. got=%T", s)
// 		return false
// 	}
// 	if letStmt.Name.Value != name {
// 		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
// 		return false
// 	}
// 	if letStmt.Name.TokenLiteral() != name {
// 		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
// 		return false
// 	}
// 	return true
// }

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
