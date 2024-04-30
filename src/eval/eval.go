package eval

import (
	"fmt"

	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/ast"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval/object"
	"bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/token"
)

func evalNode(untypedNode ast.Node) object.Object {
	switch node := untypedNode.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.StatementExpression:
		return evalNode(node.Expression)
	// Expressions
	case *ast.ExpressionBoolean:
		if node.Value {
			return TRUE
		} else {
			return FALSE
		}
	case *ast.ExpressionInteger:
		return &object.Integer{Value: node.Value}
	case *ast.ExpressionPrefix:
		return evalPrefixExpression(node.Operator, evalNode(node.Right))
	case *ast.ExpressionInfix:
		return evalInfixExpression(node.Operator, evalNode(node.Left), evalNode(node.Right))
	default:
		panic(fmt.Sprintf("unexpected case: node has type %T", node))
	}
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = evalNode(statement)
	}
	return result
}

func evalPrefixExpression(operator token.TokenType, right object.Object) object.Object {
	switch operator {
	case token.BANG:
		switch right {
		case TRUE:
			return FALSE
		case FALSE:
			return TRUE
		case NULL:
			return TRUE
		default:
			return FALSE
		}
	case token.MINUS:
		if intObj, ok := right.(*object.Integer); ok {
			return &object.Integer{Value: -intObj.Value}
		}
		return NULL
	default:
		return NULL
	}
}

func evalInfixExpression(operator token.TokenType, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.Integer).Value
		switch operator {
		case token.PLUS:
			return &object.Integer{Value: leftVal + rightVal}
		case token.MINUS:
			return &object.Integer{Value: leftVal - rightVal}
		case token.ASTERISK:
			return &object.Integer{Value: leftVal * rightVal}
		case token.SLASH:
			return &object.Integer{Value: leftVal / rightVal}
		default:
			return NULL
		}
	default:
		return NULL
	}
}
