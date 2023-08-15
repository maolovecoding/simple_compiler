package parser

import "monkey/ast"

// 普拉特语法分析器的主要思想是将解析函数（普拉特称为语义代码）与词法单元类型相关联。
// 每当遇到某个词法单元类型时，都会调用相关联的解析函数来解析对应的表达式，最后返回生成的AST节点。
type (
	prefixParseFn func() ast.Expression                          // 前缀表达式解析函数
	infixParseFn  func(expression ast.Expression) ast.Expression // 中缀表达式解析函数
	suffixParseFn func(expression ast.Expression) ast.Expression // TODO 后缀表达式解析函数
)
