package parser

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/token"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.crntToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}
func (p *Parser) parseLetStatement() *ast.LetStatement {
	s := &ast.LetStatement{Token: p.crntToken}

	if !p.advanceIfNextIs(token.IDENT) {
		return nil // could return an error here
	}

	s.Name = &ast.IdentifierExpression{Token: p.crntToken, Value: p.crntToken.Literal}

	if !p.advanceIfNextIs(token.ASSIGN) {
		return nil // could return an error here
	}

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	s := &ast.ReturnStatement{Token: p.crntToken}

	p.advance()

	// TODO: expressions
	for p.crntToken.Type != token.SEMICOLON {
		p.advance()
	}

	return s
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
