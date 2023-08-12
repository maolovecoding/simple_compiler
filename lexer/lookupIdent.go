package lexer

import "monkey/token"

var keywords = map[string]token.TokenType{
	"fn":  token.FUNCTION,
	"let": token.LET,
}

func LookupIdent(ident string) token.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // 关键字
	}
	return token.IDENT // 自定义标识符 name
}
