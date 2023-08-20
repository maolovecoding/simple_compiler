package object

// Hashable 实现改接口的对象可作为hash数据结构key 作为键必须实现该接口的
type Hashable interface {
	HashKey() HashKey
}

// TODO 去换成HashKey方法的返回值 每次生成hashKey性能消耗大
