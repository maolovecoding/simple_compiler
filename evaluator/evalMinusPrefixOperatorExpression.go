package evaluator

import "monkey/object"

// evalMinusPrefixOperatorExpression -前缀求值
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		// 非数字
		return NULL
	}
	// 对数字取反
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
