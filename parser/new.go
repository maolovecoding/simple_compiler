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
	p.registerPrefix(token.IDENT, p.parseIdentifier)       // 标识符
	p.registerPrefix(token.INT, p.parseIntegerLiteral)     // 解析INT
	p.registerPrefix(token.BANG, p.parsePrefixExpression)  // 解析前缀表达式 !
	p.registerPrefix(token.MINUS, p.parsePrefixExpression) // 解析前缀表达式 -
	// ----------------中缀表达式-------------------
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInFixExpression)
	p.registerInfix(token.MINUS, p.parseInFixExpression)
	p.registerInfix(token.SLASH, p.parseInFixExpression)
	p.registerInfix(token.ASTERISK, p.parseInFixExpression)
	p.registerInfix(token.EQ, p.parseInFixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInFixExpression)
	p.registerInfix(token.LT, p.parseInFixExpression)
	p.registerInfix(token.GT, p.parseInFixExpression)
	// ----------------中缀表达式-------------------
	// 读取两个词法单元 设置 curToken & peekToken
	p.nextToken()
	p.nextToken()
	return p
}
