package vm

import "monkey/object"

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}

// isTruthy 有值 只要不是 false 就认为是真值
func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	default:
		return true
	}
}
