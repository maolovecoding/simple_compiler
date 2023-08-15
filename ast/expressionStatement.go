package ast

import "monkey/token"

// ExpressionStatement 表达式语句
type ExpressionStatement struct {
	Token      token.Token // 表达式中的第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
