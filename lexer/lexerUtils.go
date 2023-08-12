package lexer

// lexer的一些工具方法 私有

// peekChar 查看下一个字符是什么
func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
