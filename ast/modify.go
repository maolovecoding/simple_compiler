package ast

// TODO 错误处理

type ModifierFunc func(Node) Node

// Modify 修改AST 递归遍历 符合要求的node调用修改ast的函数 然后替换原AST
func Modify(node Node, modifier ModifierFunc) Node {
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, modifier).(Expression)
	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Index, _ = Modify(node.Index, modifier).(Expression)
	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modifier).(Expression)
		node.Consequence, _ = Modify(node.Consequence, modifier).(*BlockStatement)
		node.Alternative, _ = Modify(node.Alternative, modifier).(*BlockStatement)
	case *BlockStatement:
		for i, _ := range node.Statements {
			node.Statements[i], _ = Modify(node.Statements[i], modifier).(Statement)
		}
	case *ReturnStatement:
		node.ReturnValue, _ = Modify(node.ReturnValue, modifier).(Expression)
	case *LetStatement:
		node.Value, _ = Modify(node.Value, modifier).(Expression)
	case *FunctionLiteral:
		for i, _ := range node.Parameters {
			node.Parameters[i], _ = Modify(node.Parameters[i], modifier).(*Identifier)
		}
		node.Body, _ = Modify(node.Body, modifier).(*BlockStatement)
	case *ArrayLiteral:
		for i, _ := range node.Elements {
			node.Elements[i] = Modify(node.Elements[i], modifier).(Expression)
		}
	case *HashLiteral: // 创建新的hash结构了
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey, _ := Modify(key, modifier).(Expression)
			newVal, _ := Modify(val, modifier).(Expression)
			newPairs[newKey] = newVal
		}
		node.Pairs = newPairs
	}
	return modifier(node)
}
