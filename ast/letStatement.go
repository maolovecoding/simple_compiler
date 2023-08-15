package ast

import (
	"bytes"
	"monkey/token"
)

// LetStatement let语句
type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ") // let
	out.WriteString(ls.Name.String())        // name
	out.WriteString(" = ")                   // =
	if ls.Value != nil {
		out.WriteString(ls.Value.String()) // xxx
	}
	out.WriteString(";") // let name = xxx;
	return out.String()
}
