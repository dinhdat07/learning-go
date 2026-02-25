package main

import (
	"bufio"
	"calculator/internal/engine"
	"calculator/internal/solver"
	"calculator/internal/utils"
	"fmt"
	"strings"
)

func runExpression(calculator *solver.Calculator, reader *bufio.Reader) {
	for {
		fmt.Println("\nExpression Mode")
		fmt.Println("Enter a mathematical expression to evaluate.")
		fmt.Println("Examples: 2+3*4, (1+2)/3, ans+5 (if supported)")

		expr := utils.ReadLine(reader, "> ")
		if expr == "" {
			fmt.Println("Input cannot be empty. Please enter an expression.")
			continue
		}

		ans, err := calculator.Handle(expr)
		if err != nil {
			fmt.Printf("Could not evaluate expression: %v\n", err)
			calculator.SetHasAns(false)
		} else {
			ans = utils.CleanFloat(ans)
			fmt.Printf("Result: %g\n", ans)
			calculator.SetHasAns(true)
			calculator.SetAns(ans)
		}

		if isBack := chooseNextStep(calculator, reader, ans); isBack {
			return
		}
	}
}

func runEquation(calculator *solver.Calculator, reader *bufio.Reader) {
	for {
		fmt.Println("\nEquation Mode")
		fmt.Print("\nLinear equation: a*x + b = 0")
		fmt.Print("\nQuadratic equation: a*x^2 + b*x + c = 0")
		opt := utils.ReadInt(reader, "\nChoose equation degree (1 or 2). Enter 0 to return to main menu:\n>")
		if opt == 0 {
			fmt.Println("Returning to main menu.")
			return
		}

		if opt < 0 || opt > 2 {
			fmt.Println("Invalid option. Please choose 0, 1, or 2.")
			continue
		}

		required := opt + 1
		fmt.Printf("Enter %d numbers: \n", required)

		line := utils.ReadLine(reader, "> ")

		if line == "" {
			fmt.Printf("Input cannot be empty. Please enter exactly %d numbers.", required)
			continue
		}

		nums, err := utils.ParseFloatList(line)
		if err != nil {
			fmt.Printf("Invalid number format: %v\n", err)
			fmt.Println("Please enter numbers separated by spaces")
			continue
		}

		if len(nums) != required {
			fmt.Printf("Incorrect number of values. Expected %d but received %d.\n", required, len(nums))
			continue
		}

		var ans []float64
		if opt == 1 {
			ans, err = engine.SolveLinear(nums)
		} else {
			ans, err = engine.SolveQuadratic(nums)
		}

		if err != nil {
			fmt.Printf("Failed to solve the equation: %v\n", err)
			calculator.SetHasAns(false)
		} else {
			utils.PrintSolutions(ans)
			calculator.SetHasAns(true)
		}

		if isBack := chooseNextStep(calculator, reader, ans...); isBack {
			return
		}
	}
}

func runLinearSystem(calculator *solver.Calculator, reader *bufio.Reader) {
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

				list, err := utils.ParseFloatList(line)
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

		ans, err := engine.SolveLinearSystem(matrix)
		if err != nil {
			fmt.Printf("Failed to solve the system: %v\n", err)
			fmt.Println("The system may have no solution or infinitely many solutions.")
			calculator.SetHasAns(false)
		} else {
			utils.PrintSolutions(ans)
			calculator.SetHasAns(true)
		}

		if isBack := chooseNextStep(calculator, reader, ans...); isBack {
			return
		}
	}
}

func chooseNextStep(calculator *solver.Calculator, reader *bufio.Reader, ans ...float64) (isBack bool) {
	for {
		opt := utils.ReadInt(reader, "\nWhat would you like to do next?\n1. Continue\n2. Save result to a variable\n3. Exit this mode\n> ")

		switch opt {
		case 1:
			return false

		case 2:
			if !calculator.HasAns() {
				fmt.Println("There is no valid result to save. Please compute a valid result first.")
				continue
			}

			if len(ans) == 0 {
				fmt.Println("There is no result available to save.")
				continue
			}

			chosenAns := ans[0]
			if len(ans) > 1 {
				fmt.Println("\nMultiple results detected.")
				fmt.Println("Select which result to save (1-based index).")
				for {
					index := utils.ReadInt(reader, "> ")
					if index <= 0 || index > len(ans) {
						fmt.Printf("Invalid index. Please choose a number between 1 and %d.\n", len(ans))
						continue
					}
					chosenAns = ans[index-1]
					break
				}
			}

			for {
				name := utils.ReadLine(reader, "Enter variable name (only letters, no spaces):\n> ")
				name = strings.TrimSpace(name)

				if name == "" {
					fmt.Println("Variable name cannot be empty.")
					continue
				}
				if strings.ContainsAny(name, " \t\n") {
					fmt.Println("Variable name cannot contain spaces. Please try again.")
					continue
				}

				if err := calculator.SaveVar(name, chosenAns); err != nil {
					fmt.Printf("Could not save variable: %v\n", err)
					continue
				}

				fmt.Printf("Saved '%s' = %g\n", name, chosenAns)
				return true
			}

		case 3:
			return true

		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
}
