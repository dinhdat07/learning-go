package processor

import (
	calError "calculator/error"
	. "calculator/internal"
	"calculator/internal/utils"
	"calculator/math"
	"strconv"
)

var ErrInvalidExpression = &calError.SyntaxError{
	Message: "invalid expression syntax",
}

var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'~': 3,
}

func Handle(input string) (float64, error) {
	NumStack := NewStack[float64](0)
	OpStack := NewStack[rune](0)
	prevIsValue := false

	for i := 0; i < len(input); i++ {
		c := rune(input[i])

		if c == ' ' || c == '\t' || c == '\n' {
			continue
		}

		if c == '(' {
			OpStack.Push(c)
			prevIsValue = false
			continue
		}

		if c == ')' {
			for {
				top, ok := OpStack.Peek()
				if !ok {
					return 0, ErrInvalidExpression
				}
				if top == '(' {
					break
				}
				if err := popAndCompute(NumStack, OpStack); err != nil {
					return 0, err
				}
			}

			// pop '('
			_, ok := OpStack.Pop()
			if !ok {
				return 0, ErrInvalidExpression
			}

			//check unary marker
			top, ok := OpStack.Peek()
			if ok && top == '~' {
				OpStack.Pop()
				v, okv := NumStack.Pop() // pop the num added by popAndCompute
				if !okv {
					return 0, ErrInvalidExpression
				}
				NumStack.Push(-v)
			}

			prevIsValue = true
			continue
		}

		if utils.IsOperator(c) {
			// check unary
			if (c == '+' || c == '-') && !prevIsValue {
				sign := 1.0
				for i < len(input) && (input[i] == '+' || input[i] == '-') {
					if input[i] == '-' {
						sign = -sign
					}
					i++
				}

				if i >= len(input) {
					return 0, ErrInvalidExpression
				}

				// next is number
				next := input[i]
				if (next >= '0' && next <= '9') || next == '.' {
					num, nextI, err := parseNumber(input, i)
					if err != nil {
						return 0, err
					}

					i = nextI
					NumStack.Push(sign * num)
					prevIsValue = true
					continue
				}

				// next is '('
				if next == '(' {
					//push marker '~' (unary minus)
					if sign < 0 {
						OpStack.Push('~') // unary minus marker, highest precedence
					}
					i--
					prevIsValue = false
					continue
				}

				// else -> err
				return 0, ErrInvalidExpression
			}

			// operator not unary
			for {
				top, ok := OpStack.Peek()
				if !ok || top == '(' {
					break
				}

				if hasLowerPrecedence(top, c) {
					break
				}

				if err := popAndCompute(NumStack, OpStack); err != nil {
					return 0, err
				}
			}

			OpStack.Push(c)
			prevIsValue = false
			continue
		}

		if (c >= '0' && c <= '9') || c == '.' {
			num, nextI, err := parseNumber(input, i)
			if err != nil {
				return 0, err
			}

			i = nextI
			NumStack.Push(num)
			prevIsValue = true
			continue
		}

		return 0, ErrInvalidExpression
	}

	// release stack
	for OpStack.Len() != 0 {
		top, _ := OpStack.Peek()
		if top == '~' {
			OpStack.Pop()
			v, okv := NumStack.Pop()
			if !okv {
				return 0, ErrInvalidExpression
			}
			NumStack.Push(-v)
			continue
		}

		if top == '(' {
			return 0, ErrInvalidExpression
		}
		if err := popAndCompute(NumStack, OpStack); err != nil {
			return 0, err
		}
	}

	if NumStack.Len() != 1 {
		return 0, ErrInvalidExpression
	}

	result, _ := NumStack.Peek()
	return result, nil
}

func popAndCompute(NumStack *Stack[float64], OpStack *Stack[rune]) error {
	op, okc := OpStack.Pop()
	if !okc {
		return ErrInvalidExpression
	}

	// unary minus marker: 1 operand
	if op == '~' {
		b, okb := NumStack.Pop()
		if !okb {
			return ErrInvalidExpression
		}
		NumStack.Push(-b)
		return nil
	}

	// binary operators 2 operands
	b, okb := NumStack.Pop()
	a, oka := NumStack.Pop()
	if !okb || !oka {
		return ErrInvalidExpression
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
			return err
		}
		num = ans
	default:
		return ErrInvalidExpression
	}

	NumStack.Push(num)
	return nil
}

func hasLowerPrecedence(top, curr rune) bool {
	// precedence(top) < precedence(curr): true
	return precedence[top] < precedence[curr]
}

func parseNumber(input string, i int) (float64, int, error) {
	start := i
	dot := 0

	for i < len(input) {
		ch := input[i]

		if ch == '.' {
			dot++
			if dot > 1 {
				return 0, i, ErrInvalidExpression
			}
			i++
			continue
		}

		if ch < '0' || ch > '9' {
			break
		}

		i++
	}

	numStr := input[start:i]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, i, ErrInvalidExpression
	}
	return num, i - 1, nil
}
