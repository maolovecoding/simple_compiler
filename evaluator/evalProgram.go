package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalProgram 求值
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		// return 语句 直接返回return后面的值 解包
		//if returnValue, ok := result.(*object.ReturnValue); ok {
		//	return returnValue.Value
		//}
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value // return 语句 直接返回return后面的值 解包
		case *object.Error:
			return result
		}
	}
	return result
}
