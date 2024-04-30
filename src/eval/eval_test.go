package eval

import (
	"fmt"
	"testing"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval/object"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/lexer"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/parser"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	source string
	expect object.Object
}

func TestBang(t *testing.T) {
	testCases := []testCase{
		{`!false`, TRUE},
		{`!true`, FALSE},
		{`!!false`, FALSE},
		{`!!true`, TRUE},
	}
	for i, tc := range testCases {
		doTest(t, i, tc)
	}
}

func TestBoolean(t *testing.T) {
	testCases := []testCase{
		{`false`, FALSE},
		{`true`, TRUE},
	}
	for i, tc := range testCases {
		doTest(t, i, tc)
	}
}

func TestInteger(t *testing.T) {
	testCases := []testCase{
		{`5`, &object.Integer{Value: 5}},
		{`10`, &object.Integer{Value: 10}},
	}
	for i, tc := range testCases {
		doTest(t, i, tc)
	}
}

func doTest(t *testing.T, i int, tc testCase) {
	l := lexer.New(tc.source)
	p := parser.New(l)
	program := p.ParseProgram()
	actual := evalNode(program)
	assert.Equal(t, tc.expect, actual, fmt.Sprintf("test case [%v]", i))
}
