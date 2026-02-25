package utils

import (
	calError "calculator/internal/errors"
)

var ErrInvalidFloatList = &calError.SyntaxError{
	Message: "input must be a list of valid numbers",
}
var ErrInvalidExpression = &calError.SyntaxError{
	Message: "invalid expression syntax",
}
