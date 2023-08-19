package object

// Environment 环境记录
type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	var s = map[string]Object{}
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
