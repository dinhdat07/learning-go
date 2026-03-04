package util

import (
	"fmt"
	"math"
)

func PrintSolutions(ans []float64) {
	if len(ans) == 0 {
		fmt.Println("No solution available.")
		return
	}

	if len(ans) == 1 {
		ans[0] = CleanFloat(ans[0])
		fmt.Printf("Solution:\n")
		fmt.Printf("x = %.6f\n", ans)
		return
	}

	fmt.Println("Solutions:")
	for i := range ans {
		ans[i] = CleanFloat(ans[i])
		fmt.Printf("x%d = %.6f\n", i+1, ans[i])
	}
}

func CleanFloat(v float64) float64 {
	if math.Abs(v) < 1e-9 {
		return 0
	}
	return math.Round(v*1e9) / 1e9
}
