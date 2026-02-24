package error

import "fmt"

type SyntaxError struct {
	Message string
}

func (err *SyntaxError) Error() string {
	return fmt.Sprintf("SyntaxError: %s", err.Message)
}
