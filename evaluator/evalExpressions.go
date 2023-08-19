package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalExpressions 表达式求值
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated} // 错误
		}
		result = append(result, evaluated)
	}
	return result
}
