package object

import (
	"fmt"
	"monkey/code"
)

// CompiledFunction 函数编译后的字节码表示 将其作为常量可添加到 compiler.Bytecode 然后加载到虚拟机
type CompiledFunction struct {
	Instructions code.Instructions
}

func (cf *CompiledFunction) Type() ObjectType {
	return COMPILED_FUNCTION_OBJ
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
