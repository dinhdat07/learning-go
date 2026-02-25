package main

import (
	"bufio"
	"calculator/internal/utils"
	"calculator/math"
	"calculator/processor"
	"fmt"
	"strings"
)

func runExpression(calculator *processor.Calculator, reader *bufio.Reader) {
	for {
		expr := utils.ReadLine(reader, "Please input the expression: ")
		if expr == "" {
			fmt.Println("Empty expression, try again.")
			continue
		}

		ans, err := calculator.Handle(expr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			calculator.SetHasAns(false)
		} else {
			fmt.Printf("Answer: %g\n", ans)
			calculator.SetHasAns(true)
			calculator.SetAns(ans)
		}

		isBack := chooseNextStep(calculator, reader, ans)
		if isBack {
			return
		}
	}
}

func runEquation(calculator *processor.Calculator, reader *bufio.Reader) {
	for {

		finalAns := make([]float64, 0)
		opt := utils.ReadInt(reader, "\nDegree of equation (currently support 1 - 2), 0 to back to main menu:\n>")

		switch opt {
		case 1:
			for {
				fmt.Println("\nDegree 1: a*x + b = 0")
				line := utils.ReadLine(reader, "Input a b: ")

				if line == "" {
					fmt.Println("Empty list, please input exact 2 numbers")
					continue
				}

				nums, err := processor.ParseFloatList(line)
				if err != nil {
					fmt.Println("Error", err)
					continue
				}

				if len(nums) != 2 {
					fmt.Println("Please input exact 2 numbers")
					continue
				}

				ans, err := math.SolveLinear(nums)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					calculator.SetHasAns(false)
				} else {
					fmt.Printf("Answer: %g\n", ans[0])
					calculator.SetHasAns(true)
					finalAns = ans
				}

				break
			}

		case 2:
			for {
				fmt.Println("\nDegree 2: a*x^2 + b*x + c = 0")
				line := utils.ReadLine(reader, "Input a b c: ")

				if line == "" {
					fmt.Println("Empty list, please input exact 3 numbers")
					continue
				}

				nums, err := processor.ParseFloatList(line)
				if err != nil {
					fmt.Println("Error", err)
					continue
				}

				if len(nums) != 3 {
					fmt.Println("Empty list, please input exact 3 numbers")
					continue
				}

				ans, err := math.SolveQuadratic(nums)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					calculator.SetHasAns(false)
				} else {
					fmt.Printf("Answer: %g\n", ans)
					calculator.SetHasAns(true)
					finalAns = ans
				}

				break

			}

		case 0:
			return
		default:
			fmt.Println("Not a valid option, please choose 0/1/2.")
		}

		if isBack := chooseNextStep(calculator, reader, finalAns...); isBack {
			return
		}

	}
}

func runLinearSystem(calculator *processor.Calculator, reader *bufio.Reader) {
	for {
		opt := utils.ReadInt(reader, "\nSolve Linear System\nEnter number of equations (0 to return to main menu):\n> ")

		if opt < 0 {
			fmt.Println("Invalid input. Please enter a non-negative integer (0, 1, 2, ...).")
			continue
		}

		if opt == 0 {
			fmt.Println("Returning to main menu.")
			return
		}

		rows, cols := opt, opt+1

		fmt.Printf("\nYou are solving a system with %d equations and %d variables.\n", rows, rows)
		fmt.Printf("For each equation, enter exactly %d numbers in the following format:\n", cols)
		fmt.Printf("a1 a2 ... a%d b\n", rows)
		fmt.Println("Example (3 variables): 1 1 1 6  represents 1x + 1y + 1z = 6")
		fmt.Println("Separate numbers with spaces. Decimal values are allowed (e.g., 1.5 -2 0.25).")

		matrix := make([][]float64, rows)

		for i := 0; i < rows; i++ {
			for {
				prompt := fmt.Sprintf("\nEnter equation %d of %d (%d numbers required):\n> ", i+1, rows, cols)
				line := utils.ReadLine(reader, prompt)

				if line == "" {
					fmt.Printf("Input cannot be empty. Please enter exactly %d numbers.\n", cols)
					continue
				}

				list, err := processor.ParseFloatList(line)
				if err != nil {
					fmt.Printf("Invalid number format: %v\n", err)
					fmt.Println("Please enter numbers separated by spaces, for example: 1 2 -3 4")
					continue
				}

				if len(list) != cols {
					fmt.Printf("Incorrect number of values. Expected %d numbers but received %d.\n", cols, len(list))
					fmt.Printf("Format must be: a1 a2 ... a%d b\n", rows)
					continue
				}

				matrix[i] = list
				break
			}
		}

		fmt.Println("\nSolving the system...")

		ans, err := math.SolveLinearSystem(matrix)
		if err != nil {
			fmt.Printf("Failed to solve the system: %v\n", err)
			fmt.Println("The system may have no solution or infinitely many solutions.")
			calculator.SetHasAns(false)
		} else {
			fmt.Printf("Solution:\n%g\n", ans)
			calculator.SetHasAns(true)
		}

		if isBack := chooseNextStep(calculator, reader, ans...); isBack {
			return
		}
	}
}

func chooseNextStep(calculator *processor.Calculator, reader *bufio.Reader, ans ...float64) (isBack bool) {
	for {
		opt := utils.ReadInt(reader, "\nChoose what you want to do next:\n1. Continue\n2. Save as a variable\n3. Exit this mode\n> ")

		switch opt {
		case 1:
			return false
		case 2:
			if !calculator.HasAns() {
				fmt.Println("No valid answer to save. Please calculate a valid expression first.")
				continue
			}

			chosenAns := ans[0]
			if len(ans) > 1 {
				for {
					index := utils.ReadInt(reader, "\nChoose index of the answer list (1-indexed):\n>")
					if index <= 0 || index > len(ans) {
						fmt.Println("invalid index, try again")
						continue
					}
					chosenAns = ans[index-1]
					break
				}
			}

			name := utils.ReadLine(reader, "Variable name (no spaces): ")
			name = strings.TrimSpace(name)

			if err := calculator.SaveVar(name, chosenAns); err != nil {
				fmt.Printf("Save variable failed: %v\n", err)
				continue
			}
			fmt.Println("Saved.")
			return true

		case 3:
			return true
		default:
			fmt.Println("Not a valid option, please choose 1/2/3.")
		}
	}
}
