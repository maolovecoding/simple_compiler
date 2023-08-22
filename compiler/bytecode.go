package compiler

import (
	"monkey/code"
	"monkey/object"
)

// Bytecode 字节码
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
