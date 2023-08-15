package parser

import "monkey/lexer"

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// 读取两个词法单元 设置 curToken & peekToken
	p.nextToken()
	p.nextToken()
	return p
}
