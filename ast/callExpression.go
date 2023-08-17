package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// CallExpression 调用表达式
type CallExpression struct {
	Token     token.Token  // (
	Function  Expression   // 标识符或函数字面量
	Arguments []Expression // 参数列表
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
