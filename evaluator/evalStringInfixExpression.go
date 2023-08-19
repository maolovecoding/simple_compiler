package evaluator

import "monkey/object"

// evalStringInfixExpression 字符串 操作 支持拼接拼接 TODO 字符串比较操作？
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}
