package code

import "fmt"

// Definition 操作码定义
type Definition struct {
	Name          string // 操作码名称
	OperandWidths []int  // 操作数 占多少字节 大端模式
}

var definitions = map[Opcode]*Definition{
	OpConstant:      {"OpConstant", []int{2}},      // 占用2字节 uint16 可以引用65535个常量
	OpAdd:           {"OpAdd", []int{}},            // 无操作数
	OpPop:           {"OpPop", []int{}},            // 无操作数
	OpSub:           {"OpSub", []int{}},            // 无操作数
	OpMul:           {"opMul", []int{}},            // 无操作数
	OpDiv:           {"OpDiv", []int{}},            // 无操作数
	OpTrue:          {"OpTrue", []int{}},           // 无操作数
	OpFalse:         {"OpFalse", []int{}},          // 无操作数
	OpEqual:         {"OpEqual", []int{}},          // 无操作数
	OpNotEqual:      {"OpNotEqual", []int{}},       // 无操作数
	OpGreaterThan:   {"OpGreaterThan", []int{}},    // 无操作数
	OpMinus:         {"OpMinus", []int{}},          // 无操作数
	OpBang:          {"OpBang", []int{}},           // 无操作数
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}}, // 条件跳转 栈顶元素为false跳转
	OpJump:          {"OpJump", []int{2}},          // 无条件跳转 跳转位置就是操作数所在的地址
	OpNull:          {"OpNull", []int{}},           // 压入null
	OpSetGlobal:     {"OpSetGlobal", []int{2}},     // 操作数是变量名地址
	OpGetGlobal:     {"OpGetGlobal", []int{2}},     // 操作数是变量名地址
	OpArray:         {"OpArray", []int{2}},         // 操作数是数组元素个数
	OpHash:          {"OpHash", []int{2}},          // 操作数是hash键值对个数 * 2
	OpIndex:         {"OpIndex", []int{}},
	OpCall:          {"OpCall", []int{1}},       // 函数调用指令 执行栈顶的函数 操作码是参数的个数
	OpReturnValue:   {"OpReturnValue", []int{}}, // 函数调用有返回值指令 + OpReturn 指令的能力
	OpReturn:        {"OpReturn", []int{}},      // 函数调用无返回值 为了回到调用函数的位置
	OpGetLocal:      {"OpGetLocal", []int{1}},   // 获取局部绑定的变量 操作数（地址）用1字节即可 256个局部变量够用了
	OpSetLocal:      {"OpSetLocal", []int{1}},   // 设置
	OpGetBuiltin:    {"OpGetBuiltin", []int{1}}, // 获取内置函数操作码
	OpClosure:       {"OpClosure", []int{2, 1}}, // 创建闭包 两个操作数 第一个是常量索引 去常量池找到需要转换为闭包的编译函数 第二个是栈有多少自由变量需要转移到即将创建的闭包中
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}
