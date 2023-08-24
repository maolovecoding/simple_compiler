package code

import "fmt"

// Definition 操作码定义
type Definition struct {
	Name          string // 操作码名称
	OperandWidths []int  // 操作数 占多少字节 大端模式
}

var definitions = map[Opcode]*Definition{
	OpConstant:    {"OpConstant", []int{2}},   // 占用2字节 uint16 可以引用65535个常量
	OpAdd:         {"OpAdd", []int{}},         // 无操作数
	OpPop:         {"OpPop", []int{}},         // 无操作数
	OpSub:         {"OpSub", []int{}},         // 无操作数
	OpMul:         {"opMul", []int{}},         // 无操作数
	OpDiv:         {"OpDiv", []int{}},         // 无操作数
	OpTrue:        {"OpTrue", []int{}},        // 无操作数
	OpFalse:       {"OpFalse", []int{}},       // 无操作数
	OpEqual:       {"OpEqual", []int{}},       // 无操作数
	OpNotEqual:    {"OpNotEqual", []int{}},    // 无操作数
	OpGreaterThan: {"OpGreaterThan", []int{}}, // 无操作数
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}
