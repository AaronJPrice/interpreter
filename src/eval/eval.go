package eval

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/ast"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval/object"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/parser"
)

func Evaluate(source string) (object.Object, []error) {
	program, errs := parser.Parse(source)
	if errs != nil {
		return nil, errs
	}

	return evaluateNode(program), nil
}

func evaluateNode(untypedNode ast.Node) object.Object {
	switch node := untypedNode.(type) {

	// Statements
	case *ast.Program:
		return evaluateStatements(node.Statements)

	case *ast.StatementExpression:
		return evaluateNode(node.Expression)

	// Expressions
	case *ast.ExpressionBoolean:
		if node.Value {
			return TRUE
		} else {
			return FALSE
		}

	case *ast.ExpressionInteger:
		return &object.Integer{Value: node.Value}

	default:
		panic(fmt.Sprintf("unexpected case: node has type %T", node))

	}
}

func evaluateStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = evaluateNode(statement)
	}
	return result
}
