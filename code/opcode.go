package code

// Opcode 操作码
type Opcode byte // 1字节

const (
	OpConstant Opcode = iota // 常量 常量表达式 操作码后面跟着的是操作数 用来索引常量
	OpAdd                    // 操作码 让栈顶两个元素弹栈并相加 结果入栈  操作码后不跟操作数
	OpPop                    // 操作码 弹栈 让虚拟机将栈顶元素弹出 每一个表达式语句之后都执行这个操作码
)
