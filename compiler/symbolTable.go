package compiler

// 符号表 存储和检索有关标识符 比如标识符的位置 作用域 是否被定义 绑定值的类型 以及解释过程和编译过程中有用的信息

type SymbolScope string // 作用域

const (
	GlobalScope SymbolScope = "GLOBAL"
	LocalScope  SymbolScope = "LOCAL"
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
	symbol := Symbol{Name: name, Index: s.numDefinitions}
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
		return s.Outer.Resolve(name)
	}
	return obj, ok
}
