package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// Eval 求值
func Eval(node ast.Node) object.Object {
	// 状态机 + 递归  根据 ast求值为object表示形式
	switch node := node.(type) {
	case *ast.Program:
		return evalStatement(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}
