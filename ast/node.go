package ast

type Node interface {
	TokenLiteral() string // 返回与之关联的词法单元字面量
	String() string
}
