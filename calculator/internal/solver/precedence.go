package solver

var precedence = map[rune]int{
	'+': 1, '-': 1,
	'*': 2, '/': 2,
	'~': 3,
	'^': 4,
}
