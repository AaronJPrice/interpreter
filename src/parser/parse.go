package parser

import (
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/ast"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/lexer"
)

// Convert a raw string input into an Abstract Syntax Tree
func Parse(source string) (*ast.Program, []error) {
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()
	errs := p.Errors()
	return program, errs
}
