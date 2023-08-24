package compiler

import (
	"monkey/code"
	"monkey/object"
)

// New 创建一个 compiler
func New() *Compiler {
	return &Compiler{
		instructions:        code.Instructions{},
		constants:           []object.Object{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}
}
