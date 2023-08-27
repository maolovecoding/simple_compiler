package compiler

import "monkey/code"

// CompilationScope 编译作用域
type CompilationScope struct {
	instructions        code.Instructions  // 指令集
	lastInstruction     EmittedInstruction // 最后一条指令
	previousInstruction EmittedInstruction // 倒数第二天
}
