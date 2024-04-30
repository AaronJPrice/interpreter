package lexer

import (
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"
)

// Convert a raw string input into a slice of Tokens
func Lex(source string) []token.Token {
	var ts []token.Token
	l := New(source)

	for {
		t := l.NextToken()
		if t.Type == token.EOF {
			break
		}
		ts = append(ts, t)
	}

	return ts
}
