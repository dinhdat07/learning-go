package utils

func IsOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}
