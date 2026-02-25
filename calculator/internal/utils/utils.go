package utils

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ReadLine(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := r.ReadString('\n')
	return strings.TrimSpace(line)
}

func ReadInt(r *bufio.Reader, prompt string) int {
	for {
		s := ReadLine(r, prompt)
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println("Please input a number.")
			continue
		}
		return n
	}
}

func IsOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}

func PrintSolutions(ans []float64) {
	if len(ans) == 0 {
		fmt.Println("No solution available.")
		return
	}

	if len(ans) == 1 {
		fmt.Printf("Solution:\n")
		fmt.Printf("x = %.6f\n", cleanFloat(ans[0]))
		return
	}

	fmt.Println("Solutions:")
	for i, v := range ans {
		fmt.Printf("x%d = %.6f\n", i+1, cleanFloat(v))
	}
}

func cleanFloat(v float64) float64 {
	if math.Abs(v) < 1e-9 {
		return 0
	}
	return v
}
