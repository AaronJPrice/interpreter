package parser

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	// p.registerPrefix(token.IDENT, p.parseIdentifier)

	// Avance twice to set crntToken and nextToken
	p.advance()
	p.advance()

	return p
}

type Parser struct {
	l         *lexer.Lexer
	errors    []error
	crntToken token.Token
	nextToken token.Token
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()

	for p.crntToken.Type != token.EOF {
		if s := p.parseStatement(); s != nil {
			program.Statements = append(program.Statements, s)
		}
		p.advance()
	}

	return program
}

func (p *Parser) advance() {
	p.crntToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) advanceIfNextIs(t token.TokenType) bool {
	if p.nextToken.Type == t {
		p.advance()
		return true
	} else {
		p.addError(t)
		return false
	}
}

func (p *Parser) addError(t token.TokenType) {
	err := fmt.Errorf("expected next token to be %s, got %s", t, p.nextToken.Type)
	p.errors = append(p.errors, err)
}
