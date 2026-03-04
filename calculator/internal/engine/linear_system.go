package engine

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func (e *Engine) SolveLinearSystem(matrix [][]float64) ([]float64, error) {
	if len(matrix) == 0 {
		return nil, ErrLSCannotSolved
	}
	rows, cols := len(matrix), len(matrix[0])

	A := make([]float64, 0, rows*rows)
	b := make([]float64, 0, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == rows { // last col of a row
				b = append(b, matrix[i][j])
			} else {
				A = append(A, matrix[i][j])
			}
		}
	}

	fmt.Println(A)
	fmt.Println(b)

	matA := mat.NewDense(rows, rows, A)
	matb := mat.NewDense(rows, 1, b)

	var x mat.Dense
	// solve Ax = b
	err := x.Solve(matA, matb)
	if err != nil {
		return nil, ErrLSCannotSolved
	}

	rows, _ = x.Dims()
	result := make([]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = x.At(i, 0)
	}

	return result, nil
}
