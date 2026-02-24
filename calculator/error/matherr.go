package error

import "fmt"

type MathError struct {
	Message string
}

func (err *MathError) Error() string {
	return fmt.Sprintf("MathError: %s", err.Message)
}
