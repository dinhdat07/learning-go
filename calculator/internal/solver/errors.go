package solver

import (
	calError "calculator/internal/errors"
)

var ErrInvalidExpression = &calError.SyntaxError{
	Message: "invalid expression syntax",
}

var ErrInvalidVarName = &calError.SyntaxError{
	Message: "variable name must contain only letters and no spaces",
}
