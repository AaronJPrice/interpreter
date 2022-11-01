package eval

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/ast"
	"bitbucket.org/hurricanecommerce/dev-day-2022-09-28/src/object"
)

func Evaluate(untypedNode ast.Node) object.Object {
	switch node := untypedNode.(type) {

	// Statements
	case *ast.Program:
		return evaluateStatements(node.Statements)

	case *ast.StatementExpression:
		return Evaluate(node.Expression)

		// Expressions
	case *ast.ExpressionBoolean:
		return &object.Boolean{Value: node.Value}

	case *ast.ExpressionInteger:
		return &object.Integer{Value: node.Value}

	default:
		panic(fmt.Sprintf("unexpected case: node has type %T", node))

	}
}

func evaluateStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Evaluate(statement)
	}
	return result
}
