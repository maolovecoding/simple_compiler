package ast

import (
	"bytes"
	"monkey/token"
)

type ConstStatement struct {
	Token token.Token // CONST
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

func (cs *ConstStatement) statementNode() {}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}
func (cs *ConstStatement) String() string {
	var out bytes.Buffer
	out.WriteString(cs.TokenLiteral() + " ") // const
	out.WriteString(cs.Name.String())        // name
	out.WriteString(" = ")                   // =
	if cs.Value != nil {
		out.WriteString(cs.Value.String()) // xxx
	}
	out.WriteString(";") // const name = xxx;
	return out.String()
}
