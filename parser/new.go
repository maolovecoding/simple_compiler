package parser

import (
	"monkey/lexer"
	"monkey/token"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)   // 标识符
	p.registerPrefix(token.INT, p.parseIntegerLiteral) // 解析INT
	// 读取两个词法单元 设置 curToken & peekToken
	p.nextToken()
	p.nextToken()
	return p
}
