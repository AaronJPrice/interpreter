package repl

import (
	"bufio"
	"fmt"
	"io"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/parser"
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

		program, errs := parser.Parse(scanner.Text())
		if errs != nil {
			for _, err := range errs {
				fmt.Printf("ERROR %+v\n", err)
			}
		} else {
			fmt.Printf("%+v\n", program)
		}
	}
}

func Evaluator(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		if !scanner.Scan() {
			return
		}

		l := lexer.New(scanner.Text())

		p := parser.New(l)
		program := p.ParseProgram()
		if errs := p.Errors(); errs != nil {
			for _, err := range errs {
				fmt.Printf("ERROR %+v\n", err)
			}
			continue
		}

		object := eval.Evaluate(program)
		fmt.Println(object)
	}
}
