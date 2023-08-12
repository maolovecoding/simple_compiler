package lexer

// TODO 我要如何去支持 utf-8
type Lexer struct {
	input        string
	position     int  // ch的当前位置
	readPosition int  // ch的后面一个字符位置
	ch           byte // 单个字符
}
