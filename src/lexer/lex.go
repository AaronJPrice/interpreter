package lexer

import (
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

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
