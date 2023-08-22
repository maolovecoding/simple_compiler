package vm

import (
	"fmt"
	"monkey/code"
	"monkey/object"
)

// StackSize 栈大小
const StackSize = 2048

// VM 虚拟机
type VM struct {
	constants    []object.Object   // 常量池
	instructions code.Instructions // 指令集
	stack        []object.Object   // 虚拟栈
	sp           int               // 栈指针 始终指向栈中的下一个空闲槽  栈顶的值是 stack[sp-1]
}

/*
Run 虚拟机的核心方法
实现原理就是一个主循环， 取指-解码-执行-循环
*/
func (vm *VM) Run() error {
	// 循环
	for ip := 0; ip < len(vm.instructions); ip++ {
		// 取指 栈顶指令
		op := code.Opcode(vm.instructions[ip]) // 把栈顶指令取出（字节） 转为操作码
		// 解码
		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:]) // vm.instructions[ip+1:] 操作码后面就是操作数 操作数 2字节
			ip += 2                                               // 指针加上操作数的宽度 下次循环进来指向下一个操作码
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// push 压入一个常量到虚拟栈
func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

// StackTop 获取栈顶值
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil // 空栈
	}
	return vm.stack[vm.sp-1]
}
