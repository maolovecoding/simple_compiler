package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalBlockStatement 语句块求值
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		// 发现return语句了 这里不解包 冒泡给 evalProgram 解包
		if result != nil {
			rt := result.Type()
			if rt == object.RERURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result // 直接返回了  不对后面的语句求值
			}
		}
	}
	return result
}
