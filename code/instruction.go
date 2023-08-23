package code

import (
	"bytes"
	"fmt"
)

// Instructions 指令字节的构成集合 也是指令集  每个指令都是相邻的  [指令1, 指令2, ....] [操作码1 操作数1 操作数2, ....]
type Instructions []byte

// String 每个指令都有一个string方法
func (ins Instructions) String() string {
	var out bytes.Buffer
	i := 0
	for i < len(ins) { // 指令的字节数
		def, err := Lookup(ins[i]) // 找到该指令
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:]) // 解析的是操作数 拿到操作数和操作数占的字节数
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

// fmtInstruction 格式化打印
func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths) // 操作数的数量
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}
	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0]) // 取出操作数
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}
