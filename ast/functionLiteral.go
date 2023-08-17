package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token     // fn 词法单元
	Parameters []*Identifier   // 参数列表
	Body       *BlockStatement // 函数体
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral())
	params := []string{}
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
