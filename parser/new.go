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
	// --------------true false------------------
	p.registerPrefix(token.TRUE, p.parseBoolean)  // 解析前缀表达式 -
	p.registerPrefix(token.FALSE, p.parseBoolean) // 解析前缀表达式 -
	// --------------true false------------------
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression) // 分组表达式 括号提升优先级
	p.registerPrefix(token.IF, p.parseIfExpression)          // if表达式
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral) // 解析函数字面量
	p.registerPrefix(token.STRING, p.parseStringLiteral)     // 解析字符串
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)    // 解析数组字面量
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)       // 解析 hash
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
	p.registerInfix(token.LPAREN, p.parseCallExpression)    // 调用表达式 (
	p.registerInfix(token.LBRACKET, p.parseIndexExpression) // 索引表达式 [
	// 读取两个词法单元 设置 curToken & peekToken
	p.nextToken()
	p.nextToken()
	return p
}
