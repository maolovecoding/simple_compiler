package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// evalStatement 对每一条表达式求值
func evalStatement(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}
	return result
}
