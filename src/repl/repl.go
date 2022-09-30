package repl

import (
	"bufio"
	"fmt"
	"io"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/parser"
)

const PROMPT = ">> "

func Lexer(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		if !scanner.Scan() {
			return
		}

		tokens := lexer.Lex(scanner.Text())

		for _, t := range tokens {
			fmt.Printf("%+v\n", t)
		}
	}
}

func Parser(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		if !scanner.Scan() {
			return
		}

		l := lexer.New(scanner.Text())
		p := parser.New(l)
		program := p.ParseProgram()
		fmt.Printf("%+v\n", program)
	}
}
