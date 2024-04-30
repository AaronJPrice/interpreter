package eval

import (
	"fmt"
	"testing"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/object"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/parser"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	source string
	expect object.Object
}

func TestA(t *testing.T) {
	testCases := []testCase{
		// {``, &object.Null{}},
		{`false`, &object.Boolean{Value: false}},
		{`true`, &object.Boolean{Value: true}},
		{`5`, &object.Integer{Value: 5}},
		{`10`, &object.Integer{Value: 10}},
	}
	for i, tc := range testCases {
		l := lexer.New(tc.source)
		p := parser.New(l)
		program := p.ParseProgram()
		actual := Evaluate(program)
		assert.Equal(t, tc.expect, actual, fmt.Sprintf("test case [%v]", i))
	}

}
