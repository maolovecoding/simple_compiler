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
	mainFn := &object.CompiledFunction{
		Instructions: bytecode.Instructions,
	}
	mainFrame := NewFrame(mainFn) // 主栈帧
	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame
	return &VM{
		constants: bytecode.Constants,
		//instructions: bytecode.Instructions,
		stack:       make([]object.Object, StackSize),
		sp:          0,
		globals:     make([]object.Object, GlobalsSize),
		frames:      frames,
		framesIndex: 1,
	}
}

// NewWithGlobalsStore 创建新的虚拟机对象 主要是为repl服务
func NewWithGlobalsStore(bytecode *compiler.Bytecode, s []object.Object) *VM {
	vm := New(bytecode)
	vm.globals = s
	return vm
}

// NewFrame 创建一个新的帧
func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{
		fn: fn,
		ip: -1,
	}
}
