package code

// Opcode 操作码
type Opcode byte // 1字节

const (
	OpConstant      Opcode = iota // 常量 常量表达式 操作码后面跟着的是操作数 用来索引常量
	OpAdd                         // 操作码 让栈顶两个元素弹栈并相加 结果入栈  操作码后不跟操作数
	OpPop                         // 操作码 弹栈 让虚拟机将栈顶元素弹出 每一个表达式语句之后都执行这个操作码
	OpSub                         // 减法操作码
	OpMul                         // 乘法操作码
	OpDiv                         // 除法操作码
	OpTrue                        // 压栈 true
	OpFalse                       // false
	OpEqual                       // ==
	OpNotEqual                    // !=
	OpGreaterThan                 // >  比较的是栈顶元素 < 没必要定义 通过编译器来实现代码重排序即可 比如 3 < 5 可以写成 5 > 3
	OpMinus                       // - 前缀
	OpBang                        // ! 前缀
	OpJumpNotTruthy               // 有条件跳转
	OpJump                        // 无条件跳转
	OpNull                        // 压入null
)
