package lexer

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t'
}

func isNotASpecialChar(ch rune) bool {
	return !(ch == '(' || ch == ')' || ch == ':' || ch == ';')
}
