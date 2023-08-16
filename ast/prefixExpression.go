package ast

import (
	"bytes"
	"monkey/token"
)

// PrefixExpression 前缀表达式 \n TODO 如何支持 --a ++b的？
type PrefixExpression struct {
	Token    token.Token // 前缀词法单元 比如 - !
	Operator string      // 运算符 - !
	Right    Expression  // 右侧表达式
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String() // -1 -> (-1)
}
