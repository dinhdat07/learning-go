package math

import (
	calError "calculator/error"
)

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	}
	return -1, &calError.MathError{Message: "cannot divide by zero"}

}
