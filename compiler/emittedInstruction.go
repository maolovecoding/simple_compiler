package compiler

import "monkey/code"

// EmittedInstruction 发出的指令的具体信息  操作码 & 位置(偏移量)
// 表达式语句 -> 条件表达式 -> 表达式语句 -> 块级表达式语句 -> 回到表达式语句 -> pop 导致块级表达式生成了一个pop 这个不应该生成 生成导致我们的条件判断结果会被pop掉了
//
//	如何处理？ 追踪发出的最后两条指令 包括操作码和操作码被发出的位置
type EmittedInstruction struct {
	Opcode   code.Opcode
	Position int
}
