package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		diff := (z*z - x) / (2 * z)
		if math.Abs(diff) < 1e-10 {
			break
		}
		z -= diff
	}
	return z

}

func main() {
	fmt.Println("My sqrt :", Sqrt(2))
	fmt.Println("Math sqrt:", math.Sqrt(2))
}
