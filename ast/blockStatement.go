package ast

import (
	"bytes"
	"monkey/token"
)

// BlockStatement 块级语句
type BlockStatement struct {
	Token      token.Token // { 大括号词法单元
	Statements []Statement // 语句
}

func (bs *BlockStatement) expressionNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
