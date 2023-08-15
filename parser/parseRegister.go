package parser

import "monkey/token"

// registerPrefix 给词法单元注册前缀表达式解析函数
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerPrefix 给词法单元注册中缀表达式解析函数
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
