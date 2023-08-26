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
	instructions        code.Instructions  // 指令集 字节码
	constants           []object.Object    // 常量的内在表示集 常量池
	lastInstruction     EmittedInstruction // 记录最后发出的指令
	previousInstruction EmittedInstruction // 倒数第二条发出的指令
	symbolTable         *SymbolTable       // 符号表
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
	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}
		switch node.Operator {
		case "!":
			c.emit(code.OpBang)
		case "-":
			c.emit(code.OpMinus)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}
	case *ast.IfExpression: // 表达式语句 -> 条件表达式 -> 回到表达式语句 -> pop 导致条件的值弹栈 这个不应该弹出 如何处理？ 追踪发出的最后两条指令
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}
		// TODO 先搞一个假的跳转指令 偏移量是随便写的 在编译完 node.Consequence后知道了偏移量 再修复
		jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)
		err = c.Compile(node.Consequence) // 编译 if成立的块级语句生成指令
		if err != nil {
			return err
		}
		// 如果最后一条指令是pop则移除掉 TODO 为什么？块级表达式也是表达式语句 会生成一个pop 这个不应该产生 需要干掉
		if c.lastInstructionIsPop() {
			c.removeLastPop()
		}
		// 假偏移量的 opJump
		jumpPos := c.emit(code.OpJump, 9999)
		afterConsequencePos := len(c.instructions) // 看生成块级语句的指令后 偏移量到哪里了
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)
		if node.Alternative == nil {
			c.emit(code.OpNull) // 没有else部分 生成压入null指令
		} else {
			err = c.Compile(node.Alternative)
			if err != nil {
				return err
			}
			if c.lastInstructionIsPop() {
				c.removeLastPop()
			}
		}
		afterAlternativePos := len(c.instructions)
		c.changeOperand(jumpPos, afterAlternativePos) // 回填地址
	case *ast.BlockStatement:
		var err error = nil
		for _, s := range node.Statements {
			err = c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.LetStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		symbol := c.symbolTable.Define(node.Name.Value) // 定义标识
		c.emit(code.OpSetGlobal, symbol.Index)          // 设值 栈顶的值设置到该操作数（在符号表中的地址）
	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined varible %s", node.Value) // 变量未定义 编译报错（非运行时错误）
		}
		c.emit(code.OpGetGlobal, symbol.Index)
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
	c.setLastInstruction(op, pos)
	return pos
}

// setLastInstruction 更新最后一条设置的指令 和倒数第二条
func (c *Compiler) setLastInstruction(op code.Opcode, pos int) {
	previous := c.lastInstruction
	last := EmittedInstruction{Opcode: op, Position: pos}
	c.previousInstruction = previous
	c.lastInstruction = last
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

// lastInstructionIsPop 最后一条指令是否是 pop
func (c *Compiler) lastInstructionIsPop() bool {
	return c.lastInstruction.Opcode == code.OpPop
}

// removeLastPop 删除最后一条指令
func (c *Compiler) removeLastPop() {
	c.instructions = c.instructions[:c.lastInstruction.Position]
	c.lastInstruction = c.previousInstruction
}

// replaceInstruction 指令集替换 将指令集中pos偏移位置开始的指令替换为新指令
func (c *Compiler) replaceInstruction(pos int, newInstruction []byte) {
	for i := 0; i < len(newInstruction); i++ {
		c.instructions[pos+i] = newInstruction[i]
	}
}

/*
changeOperand 改变指令集中 指定偏移量的操作数 是替换（是相同非可变长的指令）

opPos int 是指令在指令集中的位置
operand int 操作数
*/
func (c *Compiler) changeOperand(opPos int, operand int) {
	op := code.Opcode(c.instructions[opPos])
	newInstruction := code.Make(op, operand) // 根据操作码 和操作数 重新构造指令
	c.replaceInstruction(opPos, newInstruction)
}
