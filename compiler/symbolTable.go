package compiler

// 符号表 存储和检索有关标识符 比如标识符的位置 作用域 是否被定义 绑定值的类型 以及解释过程和编译过程中有用的信息

type SymbolScope string // 作用域

const (
	GlobalScope   SymbolScope = "GLOBAL"
	LocalScope    SymbolScope = "LOCAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FUNCTION" // 当对一个名称进行语法分析并使用FunctionScope返回一个符号时，我们就知道它是当前的函数名，因此它是自引用。
)

// Symbol 符号 标识
type Symbol struct {
	Name  string      // 变量名
	Scope SymbolScope // 作用域
	Index int         // 在符号表中的位置 地址
}

type SymbolTable struct {
	Outer          *SymbolTable // 父级符号表
	store          map[string]Symbol
	numDefinitions int
	FreeSymbols    []Symbol
}

// NewSymbolTable 创建符号表
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// NewEnclosedSymbolTable 基于父 SymbolTable 创建符号表
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// Define 在作用域定义一个变量 并更新其在符号表中的地址
func (s *SymbolTable) Define(name string) Symbol {
	symbol, ok := s.store[name]
	if ok {
		return symbol
	}
	symbol = Symbol{Name: name, Index: s.numDefinitions}
	if s.Outer == nil { // 全局的
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve 取变量对应的标识 symbol
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok // 没找到
		}
		// 找到了 是内置或者全部变量符号
		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}
		// 是局部变量 -> 都可以认为是自由变量(函数内用的变量不是自己定义的)
		free := s.defineFree(obj)
		return free, true // 返回自由变量 作用域是free
	}
	return obj, ok
}

// DefineBuiltin 定义内置函数符号 在 index 位置定义
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}

// 定义free变量
func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)
	symbol := Symbol{
		Name:  original.Name,
		Index: len(s.FreeSymbols) - 1,
	}
	symbol.Scope = FreeScope
	s.store[original.Name] = symbol
	return symbol
}

// DefineFunctionName 定义函数名 函数引用自身
func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{Name: name, Scope: FunctionScope, Index: 0}
	s.store[name] = symbol
	return symbol
}
