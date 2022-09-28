package lexer

import (
	"testing"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	t.Run("=+(){},;", func(t *testing.T) {
		input := `=+(){},;`
		expectedTokens := []token.Token{
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		}

		l := New(input)

		for _, expect := range expectedTokens {
			actual := l.NextToken()
			if !assert.Equal(t, expect, actual) {
				break
			}
		}
	})

	t.Run("-/*<>", func(t *testing.T) {
		input := `-/*<>`
		expectedTokens := []token.Token{
			{Type: token.MINUS, Literal: "-"},
			{Type: token.SLASH, Literal: "/"},
			{Type: token.ASTERISK, Literal: "*"},
			{Type: token.LT, Literal: "<"},
			{Type: token.GT, Literal: ">"},
		}

		l := New(input)

		for _, expect := range expectedTokens {
			actual := l.NextToken()
			if !assert.Equal(t, expect, actual) {
				break
			}
		}
	})

	t.Run("let five = 5;", func(t *testing.T) {
		input := `let five = 5;`

		expectedTokens := []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "five"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
		}

		l := New(input)

		for _, expect := range expectedTokens {
			actual := l.NextToken()
			if !assert.Equal(t, expect, actual) {
				break
			}
		}
	})

	t.Run("lots of source", func(t *testing.T) {
		input := `
		let five = 5;
		let ten = 10;

		let add = fn(x,y) {
			x + y;
		};

		let result = add(five,ten);`

		expectedTokens := []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "five"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "ten"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "10"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "add"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.IDENT, Literal: "x"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENT, Literal: "y"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.IDENT, Literal: "x"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.IDENT, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "result"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.IDENT, Literal: "add"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.IDENT, Literal: "five"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENT, Literal: "ten"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		}

		l := New(input)

		for _, expect := range expectedTokens {
			actual := l.NextToken()
			if !assert.Equal(t, expect, actual) {
				break
			}
		}
	})
}
