package lexer

import "bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"

func New(source string) *Lexer {
	l := &Lexer{source: source}
	l.advance()
	return l
}

type Lexer struct {
	source       string
	position     int
	nextPosition int
	character    byte
}

// read a token from source and advance position
func (l *Lexer) NextToken() token.Token {

	l.skipWhitespace()

	var t token.Token
	switch l.character {
	case 0:
		t = token.Token{Literal: "", Type: token.EOF}
	case '=':
		if l.peek() == '=' {
			l.advance()
			t = token.Token{Literal: "==", Type: token.EQ}
		} else {
			t = newToken(l.character, token.ASSIGN)
		}
	case '+':
		t = newToken(l.character, token.PLUS)
	case '-':
		t = newToken(l.character, token.MINUS)
	case '!':
		if l.peek() == '=' {
			l.advance()
			t = token.Token{Literal: "!=", Type: token.NOT_EQ}
		} else {
			t = newToken(l.character, token.BANG)
		}
	case '*':
		t = newToken(l.character, token.ASTERISK)
	case '/':
		t = newToken(l.character, token.SLASH)
	case '<':
		t = newToken(l.character, token.LT)
	case '>':
		t = newToken(l.character, token.GT)
	case ',':
		t = newToken(l.character, token.COMMA)
	case ';':
		t = newToken(l.character, token.SEMICOLON)
	case '(':
		t = newToken(l.character, token.LPAREN)
	case ')':
		t = newToken(l.character, token.RPAREN)
	case '{':
		t = newToken(l.character, token.LBRACE)
	case '}':
		t = newToken(l.character, token.RBRACE)
	default:
		if isLetter(l.character) {
			word := l.readWord()
			t = token.Token{Literal: word, Type: token.LookupWord(word)}
			return t
		} else if isDigit(l.character) {
			number := l.readNumber()
			t = token.Token{Literal: number, Type: token.INT}
			return t
		} else {
			t = newToken(l.character, token.ILLEGAL)
		}
	}

	l.advance()

	return t
}

func newToken(literal byte, tokenType token.TokenType) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.advance()
	}
}

func (l *Lexer) peek() byte {
	if l.nextPosition >= len(l.source) {
		return 0
	} else {
		return l.source[l.nextPosition]
	}
}

// read the next character into l.character and advance position
func (l *Lexer) advance() {
	if l.nextPosition >= len(l.source) {
		l.character = 0
	} else {
		l.character = l.source[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}

// advance position until the end of the word and return the whole thing
func (l *Lexer) readWord() string {
	return l.readString(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readString(isDigit)
}

func (l *Lexer) readString(isPartOfString func(byte) bool) string {
	start := l.position
	for isPartOfString(l.character) {
		l.advance()
	}
	return l.source[start:l.position]
}

// the character is a-z or A-Z or _
func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' ||
		'A' <= character && character <= 'Z' ||
		character == '_'
}

// the character is 0-9
func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
