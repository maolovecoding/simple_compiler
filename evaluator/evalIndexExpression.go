package evaluator

import "monkey/object"

// evalIndexExpression 索引表达式行为的求值
func evalIndexExpression(array, index object.Object) object.Object {
	switch {
	case array.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(array, index)
	default:
		return newError("index operator not supported: %s", array.Type())
	}
}

// evalArrayIndexExpression 索引表达式行为的求值
func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObj := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObj.Elements) - 1)
	if idx < 0 || idx > max {
		return NULL // 越界
	}
	return arrayObj.Elements[idx]
}
