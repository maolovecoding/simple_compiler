package vm

import (
	"monkey/compiler"
	"monkey/object"
)

/*
New 创建一个虚拟机对象
bytecode *compiler.Bytecode 字节码
*/
func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}
