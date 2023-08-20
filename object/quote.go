package object

import "monkey/ast"

// Quote 宏的一种 不对参数表达式求值
type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType {
	return QUOTE_OBJ
}
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
