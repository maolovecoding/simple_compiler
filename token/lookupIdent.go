package token

// keywords 关键字
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 找对应的标识符
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // 关键字
	}
	return IDENT // 自定义标识符 name
}
