package processor

import (
	calError "calculator/error"
	. "calculator/internal"
	"calculator/internal/utils"
	"calculator/math"
	"strconv"
	"unicode"
)

func Handle(input string) (float64, error) {
	NumStack := NewStack[float64](0)
	OpStack := NewStack[rune](0)

	for i := 0; i < len(input); i++ {
		c := rune(input[i])

		if c == ' ' || c == '\t' || c == '\n' {
			continue
		}

		if c == '(' {
			OpStack.Push(c)
			continue
		}

		if c == ')' {
			for top, _ := OpStack.Peek(); top != '('; {
				b, okb := NumStack.Pop()
				a, oka := NumStack.Pop()
				op, okc := OpStack.Pop()
				if !okb || !oka || !okc {
					return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
				}

				var num float64
				switch op {
				case '+':
					num = math.Add(a, b)
				case '-':
					num = math.Sub(a, b)
				case '*':
					num = math.Mul(a, b)
				case '/':
					ans, err := math.Div(a, b)
					if err != nil {
						return -1, err
					}
					num = ans
				default:
					return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
				}

				NumStack.Push(num)
				top, _ = OpStack.Peek()
			}

			OpStack.Pop()
			continue
		}

		if utils.IsOperator(c) {
			for {
				top, ok := OpStack.Peek()
				if !ok || top == '(' {
					break
				}

				// NOTE: precedence(top) >= precedence(c)
				topPrec := 0
				if top == '+' || top == '-' {
					topPrec = 1
				} else if top == '*' || top == '/' {
					topPrec = 2
				}
				curPrec := 0
				if c == '+' || c == '-' {
					curPrec = 1
				} else if c == '*' || c == '/' {
					curPrec = 2
				}

				if topPrec < curPrec {
					break
				}

				// pop & compute
				b, okb := NumStack.Pop()
				a, oka := NumStack.Pop()
				op, okc := OpStack.Pop()
				if !okb || !oka || !okc {
					return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
				}

				var num float64
				switch op {
				case '+':
					num = math.Add(a, b)
				case '-':
					num = math.Sub(a, b)
				case '*':
					num = math.Mul(a, b)
				case '/':
					ans, err := math.Div(a, b)
					if err != nil {
						return -1, err
					}
					num = ans
				default:
					return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
				}

				NumStack.Push(num)
			}

			OpStack.Push(c)
			continue
		}

		if unicode.IsDigit(c) || c == '.' {
			start := i
			for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.') {
				i++
			}
			numStr := input[start:i]
			i--

			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
			}
			NumStack.Push(num)
			continue
		}

		return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
	}

	for OpStack.Len() != 0 {

		b, okb := NumStack.Pop()
		a, oka := NumStack.Pop()
		op, _ := OpStack.Pop()
		if !okb || !oka {
			return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
		}

		var num float64
		switch op {
		case '+':
			num = math.Add(a, b)
		case '-':
			num = math.Sub(a, b)
		case '*':
			num = math.Mul(a, b)
		case '/':
			ans, err := math.Div(a, b)
			if err != nil {
				return -1, err
			}
			num = ans
		default:
			return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
		}

		NumStack.Push(num)
	}

	if NumStack.Len() != 1 {
		return -1, &calError.SyntaxError{Message: "your input wrong syntax"}
	}

	result, _ := NumStack.Peek()
	return result, nil
}
