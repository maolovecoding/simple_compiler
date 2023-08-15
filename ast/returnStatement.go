package ast

import "monkey/token"

type ReturnStatement struct {
	Token       token.Token // RETURN
	ReturnValue Expression  // 返回值 表达式
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
