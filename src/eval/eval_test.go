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
		{`!5`, FALSE},
		{`!!5`, TRUE},
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
		{`-5`, &object.Integer{Value: -5}},
		{`--5`, &object.Integer{Value: 5}},
	}
	for i, tc := range testCases {
		doTest(t, i, tc)
	}
}

func TestArithmeticInfixOperators(t *testing.T) {
	testCases := []testCase{
		{"5 + 5 + 5 + 5 - 10", &object.Integer{Value: 10}},
		{"2 * 2 * 2 * 2 * 2", &object.Integer{Value: 32}},
		{"-50 + 100 + -50", &object.Integer{Value: 0}},
		{"5 * 2 + 10", &object.Integer{Value: 20}},
		{"5 + 2 * 10", &object.Integer{Value: 25}},
		{"20 + 2 * -10", &object.Integer{Value: 0}},
		{"50 / 2 * 2 + 10", &object.Integer{Value: 60}},
		{"2 * (5 + 10)", &object.Integer{Value: 30}},
		{"3 * 3 * 3 + 10", &object.Integer{Value: 37}},
		{"3 * (3 * 3) + 10", &object.Integer{Value: 37}},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", &object.Integer{Value: 50}},
	}
	for i, tc := range testCases {
		doTest(t, i, tc)
	}
}
func TestBooleanInfixOperators(t *testing.T) {
	testCases := []testCase{
		{"1 < 2", TRUE},
		{"1 > 2", FALSE},
		{"1 < 1", FALSE},
		{"1 > 1", FALSE},
		{"1 == 1", TRUE},
		{"1 != 1", FALSE},
		{"1 == 2", FALSE},
		{"1 != 2", TRUE},
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
