package main

import (
	"bufio"
	"calculator/internal/utils"
	"calculator/math"
	"calculator/processor"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("======= CLI CALCULATOR ========")

	calculator := processor.NewCalculator()
	reader := bufio.NewReader(os.Stdin)

	for {
		opt := utils.ReadInt(reader, "\nChoose what you want:\n1. Expression\n2. Equation\n3. Linear System\n4. Exit\n>")

		switch opt {
		case 1:
			runExpression(calculator, reader)
		case 2:
			runEquation(calculator, reader)
		case 3:
			// runLinearSystem(calculator, reader)
		case 4:
			return
		default:
			fmt.Println("Not a valid option, please choose 1/2/3.")
		}
	}

}

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

// func runLinearSystem(calculator *processor.Calculator, reader *bufio.Reader) {

// }

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
					fmt.Println("Empty list, please input exact 2 numbers")
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

		isBack := chooseNextStep(calculator, reader, finalAns...)
		if isBack {
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
