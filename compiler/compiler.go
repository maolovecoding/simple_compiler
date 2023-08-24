package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

/*
Compiler 编译器
我们的的编译器到底应该做什么？
1. 遍历ast
2. 找到 *ast.IntegerLiteral
3. 对其进行求值 并转换为 *object.Integer
4. 将它们添加到常量字段，最后将 code.OpConstant 指令添加到内部的 code.Instructions 切片 (压栈)
*/
type Compiler struct {
	instructions code.Instructions // 指令集 字节码
	constants    []object.Object   // 常量的内在表示集 常量池
}

// Compile 编译
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		// 表达式语句产生的常量其实是临时的 使用完需要弹栈
		c.emit(code.OpPop) // 添加弹栈指令
	case *ast.InfixExpression:
		// 对 < 表达式重构为 大于表达式
		if node.Operator == "<" {
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}
			err = c.Compile(node.Left)
			if err != nil {
				return err
			}
			c.emit(code.OpGreaterThan)
			return nil
		}
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Right)
		if err != nil {
			return err
		}
		switch node.Operator {
		case "+":
			c.emit(code.OpAdd) // 发出指令
		case "-":
			c.emit(code.OpSub)
		case "*":
			c.emit(code.OpMul)
		case "/":
			c.emit(code.OpDiv)
		case ">": // 1 < 2 重构为 2 > 1
			c.emit(code.OpGreaterThan)
		case "==":
			c.emit(code.OpEqual)
		case "!=":
			c.emit(code.OpNotEqual)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}
	case *ast.IntegerLiteral:
		/*
			  思考：
			1. 对 2 的求值不会改变，始终都是2 那么 如何 ”求值“ *ast.IntegerLiteral -> *object.Integer
			2. 把 integer 添加到常量池中
		*/
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer)) // 发出指令
	case *ast.Boolean:
		if node.Value {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}

	}
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

/*
emit 生成指令并添加到最终结果（添加到文件 内存中某个区域等）

op code.Opcode 操作码

operands ...int 操作数 是地址
*/
func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins) // 添加指令 拿到指令的索引
	return pos
}

// addConstant 把求值结果添加到常量池
func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj) // 把求值结果添加到常量池
	return len(c.constants) - 1            // 返回在常量池中的索引
}

// addInstruction 添加一条新指令到指令集 并返回指令在指令集中的索引
func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return posNewInstruction
}
