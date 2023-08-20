package evaluator

import (
	"fmt"
	"monkey/ast"
	"monkey/object"
	"monkey/token"
)

// quote 不对参数表达式求值
func quote(node ast.Node, env *object.Environment) object.Object {
	// 对 unquote包裹的参数进行求值
	node = evalUnquoteCalls(node, env)
	return &object.Quote{Node: node}
}

// evalUnquoteCalls 对quoted中的unquote调用求值
func evalUnquoteCalls(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquoteCall(node) {
			return node // 不是unquote函数调用
		}
		call, ok := node.(*ast.CallExpression)
		if !ok { // 不是调用表达式
			return node
		}
		if len(call.Arguments) != 1 {
			return node // unquote函数参数也只能是1个
		}
		// 求值后得到的是Object表示 需要转为Node然后修改原ast进行替换
		unquoted := Eval(call.Arguments[0], env)
		return convertObjectToASTNode(unquoted)
	})
}

// isUnquoteCall 是否是unquote函数调用
func isUnquoteCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}
	return callExpression.Function.TokenLiteral() == "unquote"
}

// TODO 动态修改ast节点 1. 没有更新父节点的的Token字段 2. 这里是实时根据Object创建对应的Node节点 没办法包含其来源的信息 如何解决 3. 错误处理
// convertObjectToASTNode 对Object转为为Node表示
func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{
			Token: t,
			Value: obj.Value,
		}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t.Type = token.TRUE
			t.Literal = "true"
		} else {
			t.Type = token.FALSE
			t.Literal = "false"
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
