package object

import "fmt"

// 闭包对象表示
type Closure struct {
	Fn   *CompiledFunction
	Free []Object // 自由变量
}

func (c *Closure) Type() ObjectType {
	return CLOSURE_OBJ
}

func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
