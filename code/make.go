package code

import "encoding/binary"

/*
Make 编译字节码 构建指令

	op 操作码
	operands 操作数 实际上是操作数在常量池中的索引
*/
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{} // 操作码不存在
	}
	instructionLen := 1 // 指令的默认长度
	for _, w := range def.OperandWidths {
		instructionLen += w // 指令的宽度 = 操作码字节数 + 每个操作数字节数
	}
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)
	offset := 1 // 偏移量 因为操作码默认是1个字节 所以指令偏移量是1 就到了第一个操作数
	for i, o := range operands {
		width := def.OperandWidths[i] // 取出操作数宽度
		switch width {
		case 2: // 这个操作数的宽度是2 2个字节
			// 构建指令 大端模式 指令 = 操作码 + 操作数(这里o是操作数 大端模式处理为二进制然后放到指令中)
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
