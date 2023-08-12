package lexer

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
