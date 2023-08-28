package code

import "encoding/binary"

/*
ReadOperands 解码

	def 指令/操作码定义
	ins 指令[操作码, 操作数...]的操作数部分  不包含操作码

return (操作数, 操作数字节数/宽度)
*/
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths)) // 操作数
	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:])) // 转为int类型操作数了
		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))
		}
		offset += width
	}
	return operands, offset
}

// ReadUint16 二进制大端的操作数切片读取 转为无符号整形 返回
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// ReadUint8 转为无符号整形
func ReadUint8(ins Instructions) uint8 {
	return uint8(ins[0])
}
