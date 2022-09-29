package repl

import (
	"bufio"
	"fmt"
	"io"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
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

		// l := lexer.New(scanner.Text())

		// for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
		// 	fmt.Printf("%+v\n", t)
		// }
	}
}
