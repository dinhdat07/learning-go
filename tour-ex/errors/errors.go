package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("not support complex numbers, as the given number is %v", float64(e))
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}

	z := 1.0
	for {
		diff := (z*z - x) / (2 * z)
		if Abs(diff) < 1e-10 {
			break
		}
		z -= diff
	}
	return z, nil

}

func main() {
	fmt.Println(Sqrt(-2))
}
