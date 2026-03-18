package seed

import (
	"fmt"
	"math/rand"
)

func randomExpression(rng *rand.Rand) string {
	a := rng.Intn(20) + 1
	b := rng.Intn(20) + 1
	c := rng.Intn(10) + 1
	d := rng.Intn(10) + 1

	switch rng.Intn(10) {
	case 0:
		return fmt.Sprintf("%d+%d", a, b)
	case 1:
		return fmt.Sprintf("%d-%d", a, b)
	case 2:
		return fmt.Sprintf("%d*%d", a, b)
	case 3:
		den := rng.Intn(9) + 1
		num := den * (rng.Intn(10) + 1)
		return fmt.Sprintf("%d/%d", num, den)
	case 4:
		return fmt.Sprintf("%d+%d*%d", a, b, c)
	case 5:
		return fmt.Sprintf("(%d+%d)*%d", a, b, c)
	case 6:
		return fmt.Sprintf("(%d-%d)/%d", a+b, b, d)
	case 7:
		return fmt.Sprintf("%d*(%d+%d)", a, b, c)
	case 8:
		return fmt.Sprintf("(%d+%d)*(%d-%d)", a, b, c+d, d)
	default:
		// error input
		return fmt.Sprintf("%d/0", a)
	}
}

func randomLinearEquation(rng *rand.Rand) []float64 {
	a := float64(rng.Intn(19) - 9)
	b := float64(rng.Intn(41) - 20)

	// make error
	if rng.Intn(12) == 0 {
		a = 0
	}

	return []float64{a, b}
}

func randomQuadraticEquation(rng *rand.Rand) []float64 {
	a := float64(rng.Intn(9) + 1)
	b := float64(rng.Intn(41) - 20)
	c := float64(rng.Intn(41) - 20)

	switch rng.Intn(8) {
	case 0:
		// delta < 0
		a = 1
		b = 0
		c = float64(rng.Intn(10) + 1)
	case 1:
		// a=0 for invalid quadratic
		a = 0
	case 2:
		root := float64(rng.Intn(11) - 5)
		a = 1
		b = -2 * root
		c = root * root
	}

	return []float64{a, b, c}
}

func randomLinearSystem(rng *rand.Rand) [][]float64 {
	n := 2 + rng.Intn(4) // 2..5 equations

	solution := make([]float64, n)
	for i := 0; i < n; i++ {
		solution[i] = float64(rng.Intn(11) - 5)
	}

	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		row := make([]float64, n+1)

		var rhs float64
		nonZero := false

		for j := 0; j < n; j++ {
			coef := float64(rng.Intn(11) - 5)
			if coef != 0 {
				nonZero = true
			}
			row[j] = coef
			rhs += coef * solution[j]
		}

		if !nonZero {
			row[0] = 1
			rhs = solution[0]
		}

		row[n] = rhs
		matrix[i] = row
	}

	// singular / no unique solution
	if rng.Intn(10) == 0 && n >= 2 {
		copy(matrix[1], matrix[0])
	}

	return matrix
}
