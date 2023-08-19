package object

// ObjectType 值在go的表示形式 对象类型
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RERURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
)

// Object 对象接口 每个不同类型的值都有自己的表示对象形式
type Object interface {
	Type() ObjectType
	Inspect() string
}
