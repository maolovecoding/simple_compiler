package lexer

// isLetter 字符串
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit 数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
