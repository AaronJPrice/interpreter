package eval

import (
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval/object"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/parser"
)

func Evaluate(source string) (object.Object, []error) {
	program, errs := parser.Parse(source)
	if errs != nil {
		return nil, errs
	}

	return evalNode(program), nil
}
