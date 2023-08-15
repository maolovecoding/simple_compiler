package token

// TokenType TODO 类型使用int是否是更好？
type TokenType string

// 具体的 type
const (
	// 未知的词法单元
	ILLEGAL = "ILLEGAL"
	// 文件结尾
	EOF = "EOF"
	// 标识符
	IDENT = "IDENT"
	INT   = "INT"
	//	运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

/**
  TODO 1. 支持 const
       2. 支持  <= >= ==
	   3.
*/
