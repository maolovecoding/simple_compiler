package lexer

// readNumber 读取数字 返回值是数字类型的字符串 TODO  [1]=> 123 [2]=> 12.01 [3] => 12_311_11
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

// TODO 1. 只能支持整形 2. 如何支持浮点数？ 3. 如何支持便于记忆的下划线数字？
