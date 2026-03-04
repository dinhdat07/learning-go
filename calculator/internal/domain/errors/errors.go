package error

import "fmt"

type MathError struct {
	Message string
}

func (err *MathError) Error() string {
	return fmt.Sprintf("MathError: %s", err.Message)
}

type SyntaxError struct {
	Message string
}

func (err *SyntaxError) Error() string {
	return fmt.Sprintf("SyntaxError: %s", err.Message)
}
