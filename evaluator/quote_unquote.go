package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// quote 不对参数表达式求值
func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
