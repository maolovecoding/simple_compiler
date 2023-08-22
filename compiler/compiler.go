package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

// Compiler 编译器
type Compiler struct {
	instructions code.Instructions // 指令集 字节码
	constants    []object.Object   // 常量的内在表示集 常量池
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{}
}
