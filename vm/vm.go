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
		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv: // 1 + 2
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return nil
			}
		case code.OpEqual, code.OpNotEqual, code.OpGreaterThan:
			err := vm.executeComparison(op)
			if err != nil {
				return nil
			}
		case code.OpBang:
			err := vm.executeBangOperator()
			if err != nil {
				return nil
			}
		case code.OpMinus:
			err := vm.executeMinusOperator()
			if err != nil {
				return nil
			}
		case code.OpTrue:
			err := vm.push(True)
			if err != nil {
				return err
			}
		case code.OpFalse:
			err := vm.push(False)
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		}
	}
	return nil
}

// executeBinaryOperation 二元运算 + - * /
func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop() // 2
	left := vm.pop()  // 1
	// 操作数类型是否兼容
	leftType := left.Type()
	rightType := right.Type()
	if leftType == object.INTEGER_OBJ && rightType == object.INTEGER_OBJ {
		return vm.executeBinaryIntegerOperation(op, left, right)
	}
	return fmt.Errorf("unsupported types for binary operation: %s %s", leftType, rightType)
}

// executeBinaryIntegerOperation integer运算
func (vm *VM) executeBinaryIntegerOperation(op code.Opcode, left, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value
	var result int64
	switch op {
	case code.OpAdd:
		result = leftValue + rightValue
	case code.OpSub:
		result = leftValue - rightValue
	case code.OpMul:
		result = leftValue * rightValue
	case code.OpDiv:
		result = leftValue / rightValue
	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}
	return vm.push(&object.Integer{Value: result})
}

// executeComparison 比较运算
func (vm *VM) executeComparison(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()
	if right.Type() == object.INTEGER_OBJ && left.Type() == object.INTEGER_OBJ {
		return vm.executeIntegerComparison(op, left, right)
	}
	switch op {
	case code.OpEqual:
		return vm.push(nativeBoolToBooleanObject(right == left))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObject(right != left))
	default:
		return fmt.Errorf("unknown operator: %d (%s %s)", op, left.Type(), right.Type())
	}
}

// executeIntegerComparison 整数比较
func (vm *VM) executeIntegerComparison(op code.Opcode, left, right object.Object) error {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch op {
	case code.OpEqual:
		return vm.push(nativeBoolToBooleanObject(leftVal == rightVal))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObject(leftVal != rightVal))
	case code.OpGreaterThan:
		return vm.push(nativeBoolToBooleanObject(leftVal > rightVal))
	default:
		return fmt.Errorf("unknown operator: %d", op)
	}
}

// executeBangOperator 取反运算符
func (vm *VM) executeBangOperator() error {
	operand := vm.pop()
	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	default:
		return vm.push(False)
	}
}

// executeMinusOperator -前缀运算
func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()
	if operand.Type() != object.INTEGER_OBJ {
		return fmt.Errorf("unsupported type for negation: %s", operand.Type())
	}
	value := operand.(*object.Integer).Value
	return vm.push(&object.Integer{Value: -value})
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

/*
pop 弹栈
弹栈，其实就是取出栈顶元素，然后栈指针下移（下次入栈覆盖了这些弹出的栈顶元素）
我们所谓的内存擦除等等操作，其实就是释放内存的使用权，并不是说我们释放内存的时候对内存的数据进行清零操作
*/
func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

// StackTop 获取栈顶值
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil // 空栈
	}
	return vm.stack[vm.sp-1]
}

// LastPoppedStackElem 测试方法 在 code.OpPop 执行之前 栈顶元素不应该发送改变
// 因为我们执行了 pop 已经弹栈了 但是我们栈没有清空 这里是为了验证栈顶元素的正确性
func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}
