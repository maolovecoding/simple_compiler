package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalProgram 求值
func evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement)
		// return 语句 直接返回return后面的值 解包
		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}
	return result
}
