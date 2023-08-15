package ast

import (
	"monkey/token"
)

// Identifier 标识符
type Identifier struct {
	Token token.Token // IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
