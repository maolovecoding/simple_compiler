package object

// Environment 环境记录
type Environment struct {
	store map[string]Object
	outer *Environment // 父环境
}

func NewEnvironment() *Environment {
	var s = map[string]Object{}
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// NewEnclosedEnvironment 创建新的环境 函数级别 TODO 块级别
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
