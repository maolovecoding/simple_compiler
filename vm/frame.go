package vm

import (
	"monkey/code"
	"monkey/object"
)

const MaxFrames = 1024 // 最大帧数

// Frame 调用函数时需要需要记录的信息 也就是帧[栈]对象 记录编译函数 和栈帧
type Frame struct {
	fn *object.CompiledFunction // 编译的函数
	ip int                      // 栈帧
}

// Instructions 获取函数指令集
func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
