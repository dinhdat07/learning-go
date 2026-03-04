package cli

import (
	"calculator/internal/model"
	util "calculator/internal/util"
	"fmt"
)

func (app *App) runExpression() {
	for {
		fmt.Println("\nExpression Mode")
		fmt.Println("Enter a mathematical expression to evaluate.")
		fmt.Println("Examples: 2+3*4, (1+2)/3, ans+5 (if supported)")

		expr := util.ReadLine(app.reader, "> ")
		if expr == "" {
			fmt.Println("Input cannot be empty. Please enter an expression.")
			continue
		}

		ans, err, durationMs := app.calculatorService.EvalExpression(expr)
		app.historyService.Save(model.ExpressionMode, expr, ans, err, durationMs)

		if isBack := app.chooseNextStep(ans); isBack {
			return
		}
	}
}

func (app *App) runEquation() {
	for {
		fmt.Println("\nEquation Mode")
		fmt.Print("\nLinear equation: a*x + b = 0")
		fmt.Print("\nQuadratic equation: a*x^2 + b*x + c = 0")
		opt := util.ReadInt(app.reader, "\nChoose equation degree (1 or 2). Enter 0 to return to main menu:\n>")
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

		line := util.ReadLine(app.reader, "> ")

		if line == "" {
			fmt.Printf("Input cannot be empty. Please enter exactly %d numbers.", required)
			continue
		}

		nums, err := util.ParseFloatList(line)
		if err != nil {
			fmt.Printf("Invalid number format: %v\n", err)
			fmt.Println("Please enter numbers separated by spaces")
			continue
		}

		if len(nums) != required {
			fmt.Printf("Incorrect number of values. Expected %d but received %d.\n", required, len(nums))
			continue
		}

		ans, err, duration := app.calculatorService.SolveEquation(opt, nums)
		app.historyService.Save(model.EquationMode, line, ans, err, duration)

		if isBack := app.chooseNextStep(ans...); isBack {
			return
		}
	}
}

func (app *App) runLinearSystem() {
	for {
		opt := util.ReadInt(app.reader, "\nSolve Linear System\nEnter number of equations (0 to return to main menu):\n> ")

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
				line := util.ReadLine(app.reader, prompt)

				if line == "" {
					fmt.Printf("Input cannot be empty. Please enter exactly %d numbers.\n", cols)
					continue
				}

				list, err := util.ParseFloatList(line)
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

		ans, err, duration := app.calculatorService.SolveLinearSystem(matrix)
		app.historyService.Save(model.LinearSystemMode, matrix, ans, err, duration)

		if isBack := app.chooseNextStep(ans...); isBack {
			return
		}
	}
}

func (app *App) chooseNextStep(ans ...float64) (isBack bool) {
	for {
		opt := util.ReadInt(app.reader, "\nWhat would you like to do next?\n1. Continue\n2. Save result to a variable\n3. Exit this mode\n> ")

		switch opt {
		case 1:
			return false

		case 2:
			return app.calculatorService.SaveAnsToVar(ans, app.reader)
		case 3:
			return true

		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
}
