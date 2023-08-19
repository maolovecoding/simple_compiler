package lexer

func (l *Lexer) readString() string {
	position := l.position + 1 // 跳过 "
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
