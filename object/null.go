package object

type Null struct {
}

func (n *Null) Inspect() string {
	return "NULL"
}

func (n *Null) Type() ObjectType {
	return INTEGER_OBJ
}
