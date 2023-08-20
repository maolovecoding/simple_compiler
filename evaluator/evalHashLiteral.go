package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalHashLiteral 对hash字面量求值
func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, keyValue := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}
		value := Eval(keyValue, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey() // 对key求hash
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}
	return &object.Hash{Pairs: pairs}
}
