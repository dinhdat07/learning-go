package processor

import (
	calError "calculator/error"
	. "calculator/internal"
	"calculator/internal/utils"
)

var ErrInvalidExpression = &calError.SyntaxError{
	Message: "invalid expression syntax",
}

var ErrInvalidVarName = &calError.SyntaxError{
	Message: "variable name must contain only letters (a-z or A-Z) and no spaces",
}

var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'~': 3,
	'^': 4,
}

type Calculator struct {
	hasAns    bool // to check if the prev equation/expression have valid result
	ans       float64
	preAns    float64
	variables map[string]float64
}

func NewCalculator() *Calculator {
	return &Calculator{
		hasAns:    false,
		ans:       0,
		preAns:    0,
		variables: make(map[string]float64),
	}
}

func (c *Calculator) SaveVar(name string, value float64) error {
	if !isValidVarName(name) {
		return ErrInvalidVarName
	}
	c.variables[name] = value
	return nil
}

func (c *Calculator) Ans() float64 {
	return c.ans
}

func (c *Calculator) PreAns() float64 {
	return c.preAns
}

func (c *Calculator) SetAns(ans float64) {
	c.preAns, c.ans = c.ans, ans
}

func (c *Calculator) HasAns() bool {
	return c.hasAns
}

func (c *Calculator) SetHasAns(hasAns bool) {
	c.hasAns = hasAns
}

func (calculator *Calculator) Handle(input string) (float64, error) {
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
					NumStack.Push(num)
					if sign < 0 {
						OpStack.Push('~')
					}
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

				if !shouldPop(top, c) {
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

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			num, nextI, err := parseKeyword(calculator, input, i)
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
	calculator.preAns = calculator.ans
	calculator.ans = result
	return result, nil
}
