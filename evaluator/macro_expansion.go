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
