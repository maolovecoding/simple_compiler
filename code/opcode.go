package code

// Opcode 操作码
type Opcode byte // 1字节

const (
	OpConstant Opcode = iota // 常量 常量表达式 操作码后面跟着的是操作数 用来索引常量
)
