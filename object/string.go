package object

import "hash/fnv"

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

// HashKey 生成hash值 TODO 如何解决hash碰撞 单链法  开放寻址法
func (s *String) HashKey() HashKey {
	h := fnv.New64a()                                // 创建一个64位的hash对象
	h.Write([]byte(s.Value))                         // 添加计算hash值的数据
	return HashKey{Type: s.Type(), Value: h.Sum64()} // Sum64 拿到计算后的hash值
}
