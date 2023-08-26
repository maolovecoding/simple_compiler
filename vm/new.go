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
		globals:      make([]object.Object, GlobalsSize),
	}
}

// NewWithGlobalsStore 创建新的虚拟机对象 主要是为repl服务
func NewWithGlobalsStore(bytecode *compiler.Bytecode, s []object.Object) *VM {
	vm := New(bytecode)
	vm.globals = s
	return vm
}
