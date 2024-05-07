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
		return nativeBoolToBooleanObject(node.Value)
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
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left.(*object.Integer), right.(*object.Integer))
		// } else if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
		// 	return evalBooleanInfixExpression(operator, left.(*object.Boolean), right.(*object.Boolean))
	} else {
		return NULL
	}
}

func evalIntegerInfixExpression(operator token.TokenType, left, right *object.Integer) object.Object {
	switch operator {
	case token.PLUS:
		return &object.Integer{Value: left.Value + right.Value}
	case token.MINUS:
		return &object.Integer{Value: left.Value - right.Value}
	case token.ASTERISK:
		return &object.Integer{Value: left.Value * right.Value}
	case token.SLASH:
		return &object.Integer{Value: left.Value / right.Value}
	case token.LT:
		return nativeBoolToBooleanObject(left.Value < right.Value)
	case token.GT:
		return nativeBoolToBooleanObject(left.Value > right.Value)
	case token.EQ:
		return nativeBoolToBooleanObject(left.Value == right.Value)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(left.Value != right.Value)
	default:
		return NULL
	}
}

// func evalBooleanInfixExpression(operator token.TokenType, left, right *object.Boolean) object.Object {
// 	switch operator {
// 	case token.EQ:
// 		return nativeBoolToBooleanObject(left == right)
// 	case token.NOT_EQ:
// 		return nativeBoolToBooleanObject(left != right)
// 	default:
// 		return NULL
// 	}
// }
