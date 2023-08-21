package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// DefineMacros 查找宏定义 并从AST中删除
func DefineMacros(program *ast.Program, env *object.Environment) {
	var definitions []int
	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)             // 添加宏定义到环境中
			definitions = append(definitions, i) // 跟踪宏定义的位置
		}
	}
	// 删除宏定义语句
	for i := len(definitions) - 1; i >= 0; i-- {
		definitionIndex := definitions[i]
		program.Statements = append(
			program.Statements[:definitionIndex],
			program.Statements[definitionIndex+1:]...,
		)
	}
}

// ExpandMacros 展开宏定义
func ExpandMacros(program ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		callExpression, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}
		macro, ok := isMacroCall(callExpression, env)
		if !ok {
			return node
		}
		args := quoteArgs(callExpression)
		evalEnv := extendMacroEnv(macro, args)
		evaluated := Eval(macro.Body, evalEnv)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes form macros")
		}
		return quote.Node
	})
}

// addMacro 将宏定义语句添加到环境中
func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)
	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}
	env.Set(letStatement.Name.Value, macro)
}

// isMacroDefinition 是否是宏定义语句的ast
func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}
	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}
	return true
}

// isMacroCall 宏表达式调用
func isMacroCall(exp *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	identifier, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}
	obj, ok := env.Get(identifier.Value) // 取出标识符 也就是宏定义函数
	if !ok {
		return nil, false
	}
	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}
	return macro, true
}

// quoteArgs 参数转为 Quote形式表示
func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	var args []*object.Quote
	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}
	return args
}

// extendMacroEnv 创建macro函数的环境
func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)
	for idx, param := range macro.Parameters {
		extended.Set(param.Value, args[idx])
	}
	return extended
}
