package token

// Token 词法单元 类型 + 字面量
type Token struct {
	Type    TokenType
	Literal string
}
