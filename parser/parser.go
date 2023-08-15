package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser 语法解析
type Parser struct {
	l              *lexer.Lexer                      // 词法分析
	curToken       token.Token                       // 当前token
	peekToken      token.Token                       // 下一个token
	errors         []string                          // 错误
	prefixParseFns map[token.TokenType]prefixParseFn // 解析函数
	infixParseFns  map[token.TokenType]infixParseFn  // 解析函数
}

// nextToken 下一个token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram 解析整个程序生成 ast
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// Errors 所有错误
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError 添加错误
func (p *Parser) peekError(tokenType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tokenType, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
