package ast

import "monkey/token"

// IntegerLiteral 整形字面量
type IntegerLiteral struct {
	Token token.Token
	Value int64 // "5" -> 5
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
