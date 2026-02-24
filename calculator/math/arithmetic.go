package math

import (
	calError "calculator/error"
	gomath "math"
)

var ErrDivideByZero = &calError.MathError{
	Message: "division by zero",
}

var ErrInvalidPower = &calError.MathError{
	Message: "invalid power operation",
}

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
	return 0, ErrDivideByZero

}

func Pow(a, b float64) (float64, error) {
	if a == 0 && b == 0 {
		return 0, ErrInvalidPower
	}
	result := gomath.Pow(a, b)

	if gomath.IsNaN(result) || gomath.IsInf(result, 0) {
		return 0, ErrInvalidPower
	}

	return result, nil
}
