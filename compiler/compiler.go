package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/object"
	"sort"
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
	//instructions        code.Instructions  // 指令集 字节码
	constants []object.Object // 常量的内在表示集 常量池
	//lastInstruction     EmittedInstruction // 记录最后发出的指令
	//previousInstruction EmittedInstruction // 倒数第二条发出的指令
	symbolTable *SymbolTable       // 符号表
	scopes      []CompilationScope // 编译作用域
	scopeIndex  int                // 在那个作用域
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
		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}
		// 假偏移量的 opJump
		jumpPos := c.emit(code.OpJump, 9999)
		afterConsequencePos := len(c.currentInstructions()) // 看生成块级语句的指令后 偏移量到哪里了
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)
		if node.Alternative == nil {
			c.emit(code.OpNull) // 没有else部分 生成压入null指令
		} else {
			err = c.Compile(node.Alternative)
			if err != nil {
				return err
			}
			if c.lastInstructionIs(code.OpPop) {
				c.removeLastPop()
			}
		}
		afterAlternativePos := len(c.currentInstructions())
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
		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index) // 设值 栈顶的值设置到该操作数（在符号表中的地址）
		} else {
			c.emit(code.OpSetLocal, symbol.Index) // 局部符号表
		}
	case *ast.ArrayLiteral:
		for _, ele := range node.Elements {
			err := c.Compile(ele)
			if err != nil {
				return err
			}
		}
		c.emit(code.OpArray, len(node.Elements))
	case *ast.HashLiteral:
		keys := []ast.Expression{}
		for k := range node.Pairs {
			keys = append(keys, k)
		}
		// go中每次遍历map的键值对不能保证键的一致顺序 手动排序 保证生成的指令一致
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
		for _, k := range keys {
			err := c.Compile(k)
			if err != nil {
				return err
			}
			err = c.Compile(node.Pairs[k])
			if err != nil {
				return err
			}
		}
		c.emit(code.OpHash, len(node.Pairs)*2) // 键值对的两倍 编译前计算好
	case *ast.IndexExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Index)
		if err != nil {
			return err
		}
		c.emit(code.OpIndex)
	case *ast.FunctionLiteral:
		c.enterScope() // 进入新的编译作用域 函数指令在这个作用域下生成
		err := c.Compile(node.Body)
		if err != nil {
			return err
		}
		if c.lastInstructionIs(code.OpPop) { // 函数最后一条指令是pop 隐式函数返回值 把指令转为 OpReturnValue
			c.replaceLastPopWithReturn()
		}
		if !c.lastInstructionIs(code.OpReturnValue) {
			// 空函数体 fn () {} & 不能转换为该语句的情况 比如 let name = "zs";
			c.emit(code.OpReturn)
		}
		instructions := c.leaveScope() // 函数作用域下生成的指令集
		compiledFn := &object.CompiledFunction{Instructions: instructions}
		c.emit(code.OpConstant, c.addConstant(compiledFn)) // 编译函数字面量 添加到常量池
	case *ast.ReturnStatement:
		err := c.Compile(node.ReturnValue)
		if err != nil {
			return err
		}
		c.emit(code.OpReturnValue)
	case *ast.CallExpression:
		err := c.Compile(node.Function)
		if err != nil {
			return err
		}
		c.emit(code.OpCall)
	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined varible %s", node.Value) // 变量未定义 编译报错（非运行时错误）
		}
		if symbol.Scope == GlobalScope {
			c.emit(code.OpGetGlobal, symbol.Index)
		} else {
			c.emit(code.OpGetLocal, symbol.Index)
		}
	case *ast.StringLiteral:
		str := &object.String{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(str))
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
		Instructions: c.currentInstructions(),
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
	previous := c.scopes[c.scopeIndex].lastInstruction
	last := EmittedInstruction{Opcode: op, Position: pos}
	c.scopes[c.scopeIndex].previousInstruction = previous
	c.scopes[c.scopeIndex].lastInstruction = last
}

// addConstant 把求值结果添加到常量池
func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj) // 把求值结果添加到常量池
	return len(c.constants) - 1            // 返回在常量池中的索引
}

// addInstruction 添加一条新指令到指令集 并返回指令在指令集中的索引
func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.currentInstructions())
	updatedInstructions := append(c.currentInstructions(), ins...)
	c.scopes[c.scopeIndex].instructions = updatedInstructions
	return posNewInstruction
}

// lastInstructionIs 断言最后一条指令是否是 op
func (c *Compiler) lastInstructionIs(op code.Opcode) bool {
	if len(c.currentInstructions()) == 0 {
		return false
	}
	return c.scopes[c.scopeIndex].lastInstruction.Opcode == op
}

// removeLastPop 删除最后一条指令
func (c *Compiler) removeLastPop() {
	last := c.scopes[c.scopeIndex].lastInstruction
	previous := c.scopes[c.scopeIndex].previousInstruction
	old := c.currentInstructions()
	newIns := old[:last.Position]
	c.scopes[c.scopeIndex].instructions = newIns
	c.scopes[c.scopeIndex].lastInstruction = previous
}

// replaceInstruction 指令集替换 将指令集中pos偏移位置开始的指令替换为新指令
func (c *Compiler) replaceInstruction(pos int, newInstruction []byte) {
	ins := c.currentInstructions()
	for i := 0; i < len(newInstruction); i++ {
		ins[pos+i] = newInstruction[i]
	}
}

/*
changeOperand 改变指令集中 指定偏移量的操作数 是替换（是相同非可变长的指令）

opPos int 是指令在指令集中的位置
operand int 操作数
*/
func (c *Compiler) changeOperand(opPos int, operand int) {
	op := code.Opcode(c.scopes[c.scopeIndex].instructions[opPos])
	newInstruction := code.Make(op, operand) // 根据操作码 和操作数 重新构造指令
	c.replaceInstruction(opPos, newInstruction)
}

// currentInstructions 获取当前作用域的指令集
func (c *Compiler) currentInstructions() code.Instructions {
	return c.scopes[c.scopeIndex].instructions
}

// enterScope 进入一个新的编译作用域 生成的指令在该作用域下
func (c *Compiler) enterScope() {
	scope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable) // 函数作用域下的符号表
	c.scopeIndex++
}

// leaveScope 离开当前编译作用域 回到父作用域
func (c *Compiler) leaveScope() code.Instructions {
	instructions := c.currentInstructions()
	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--
	c.symbolTable = c.symbolTable.Outer // 回到调用方作用域下的符号表
	return instructions
}

// replaceLastPopWithReturn 函数返回值隐式 最后一条指令是pop转换为 returnValue 指令是等长的 注意
func (c *Compiler) replaceLastPopWithReturn() {
	lastPos := c.scopes[c.scopeIndex].lastInstruction.Position
	c.replaceInstruction(lastPos, code.Make(code.OpReturnValue))       // 只有操作码
	c.scopes[c.scopeIndex].lastInstruction.Opcode = code.OpReturnValue // 替换操作码
}
