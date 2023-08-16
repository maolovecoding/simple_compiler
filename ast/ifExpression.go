package ast

import (
	"bytes"
	"monkey/token"
)

type IfExpression struct {
	Token       token.Token     // if 词法单元
	Condition   Expression      // 条件表达式
	Consequence *BlockStatement // 结果 块表达式 代码块
	Alternative *BlockStatement // 替代结果 块表达式
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String()) // if 代码块
	// else代码块
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
