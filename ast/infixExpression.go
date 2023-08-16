package ast

import (
	"bytes"
	"monkey/token"
)

// InfixExpression 中缀表达式

type InfixExpression struct {
	Token    token.Token // 运算符的词法单元 + - *
	Left     Expression
	Operator string     // 运算符 + - 字面量
	Right    Expression // 右侧表达式
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
