package parser

import "monkey/token"

// peekTokenIs 偷看下一个token是否是目标类型
func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

// curTokenIs 当前token是否是目标类型
func (p *Parser) curTokenIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

// expectPeek 下一个token是期待token 则断言成功 吃掉当前的token 后移
func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}
	p.peekError(tokenType) // 添加错误
	return false
}
