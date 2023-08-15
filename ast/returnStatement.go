package ast

import (
	"bytes"
	"monkey/token"
)

type ReturnStatement struct {
	Token       token.Token // RETURN
	ReturnValue Expression  // 返回值 表达式
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ") // return
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";") // return xxx;
	return out.String()
}
