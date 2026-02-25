package engine

import (
	calError "calculator/internal/errors"
)

var ErrDivideByZero = &calError.MathError{
	Message: "division by zero is not allowed",
}

var ErrInvalidPower = &calError.MathError{
	Message: "invalid exponent operation",
}

var ErrLSCannotSolved = &calError.MathError{
	Message: "the system does not have a unique solution",
}

var ErrInvalidInput = &calError.SyntaxError{
	Message: "incorrect number of values in input",
}
var ErrNoSolution = &calError.MathError{
	Message: "the equation has no solution",
}
