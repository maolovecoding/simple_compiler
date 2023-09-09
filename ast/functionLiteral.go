package ast

import (
	"bytes"
	"fmt"
	"monkey/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token     // fn 词法单元
	Parameters []*Identifier   // 参数列表
	Body       *BlockStatement // 函数体
	Name       string          // 捕获函数名
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral())
	if fl.Name != "" {
		out.WriteString(fmt.Sprintf("<%s>", fl.Name))
	}
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
