package ast

import "monkey/token"

type ConstStatement struct {
	Token token.Token // CONST
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

func (cs *ConstStatement) statementNode() {}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}
