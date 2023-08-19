package evaluator

import (
	"monkey/object"
)

// applyFunction 函数调用
func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}
	extendsEnv := extendFunctionEnv(function, args)
	evaluated := Eval(function.Body, extendsEnv)
	return unwrapReturnValue(evaluated)
}

// extendFunctionEnv 根据父env 创建函数自己的env
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for parentIndex, param := range fn.Parameters {
		env.Set(param.Value, args[parentIndex]) // 添加环境记录 变量
	}
	return env
}

// unwrapReturnValue 解包 如果是return 就返回return的内容
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}
