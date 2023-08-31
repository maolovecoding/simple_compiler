package compiler

import (
	"monkey/code"
	"monkey/object"
)

// New 创建一个 compiler
func New() *Compiler {
	mainScope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}
	symbolTable := NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name) // 定义内置函数
	}
	return &Compiler{
		//instructions:        code.Instructions{},
		constants: []object.Object{},
		//lastInstruction:     EmittedInstruction{},
		//previousInstruction: EmittedInstruction{},
		symbolTable: symbolTable,
		scopes:      []CompilationScope{mainScope},
		scopeIndex:  0,
	}
}

// NewWithState 创建编译器对象 覆盖已有的符号表和常量
func NewWithState(s *SymbolTable, constants []object.Object) *Compiler {
	compiler := New()
	compiler.symbolTable = s
	compiler.constants = constants
	return compiler
}
