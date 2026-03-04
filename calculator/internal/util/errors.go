package util

import (
	calError "calculator/internal/domain/errors"
)

var ErrInvalidFloatList = &calError.SyntaxError{
	Message: "input must be a list of valid numbers",
}
