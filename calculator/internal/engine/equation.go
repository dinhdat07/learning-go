package engine

import (
	"math"
)

func SolveLinear(nums []float64) ([]float64, error) {
	// ax + b = 0
	// x = -b/a
	if len(nums) != 2 {
		return nil, ErrInvalidInput
	}

	a := nums[0]
	b := nums[1]

	res, err := Div(-b, a)
	if err != nil {
		return nil, ErrNoSolution
	}
	return []float64{res}, nil
}

func SolveQuadratic(nums []float64) ([]float64, error) {
	// ax^2 + b^x + c = 0
	// delta = b^2 - 4 *a*c

	if len(nums) != 3 {
		return nil, ErrInvalidInput
	}

	a, b, c := nums[0], nums[1], nums[2]
	if a == 0 {
		return SolveLinear(nums[1:3])
	}

	delta := b*b - 4*a*c
	switch {
	case delta > 0:
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		return []float64{x1, x2}, nil
	case delta == 0:
		x := -b / (2 * a)
		return []float64{x}, nil
	case delta < 0:
		return nil, ErrNoSolution
	}

	return nil, ErrInvalidInput
}
