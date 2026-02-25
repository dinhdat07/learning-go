package processor

import (
	"calculator/internal/stack"
	"calculator/math"
)

func popAndCompute(NumStack *stack.Stack[float64], OpStack *stack.Stack[rune]) error {
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
	case '^':
		ans, err := math.Pow(a, b)
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

func shouldPop(top, curr rune) bool {
	// right-associative
	if curr == '^' {
		return precedence[top] > precedence[curr]
	}

	// left-associative
	return precedence[top] >= precedence[curr]
}

func isValidVarName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, ch := range name {
		if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
			return false
		}
	}
	return true
}
