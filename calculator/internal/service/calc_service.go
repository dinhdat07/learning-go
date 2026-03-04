package service

import (
	"bufio"
	"calculator/internal/domain/solver"
	util "calculator/internal/util"
	"fmt"
	"strings"
	"time"
)

type CalculatorService struct {
	calculator *solver.Calculator
}

func NewCalculatorService() *CalculatorService {
	calculator := solver.NewCalculator()
	return &CalculatorService{calculator}
}

func (svc *CalculatorService) EvalExpression(expr string) (float64, error, int64) {
	start := time.Now()
	ans, err := svc.calculator.Handle(expr)
	duration := time.Since(start)
	durationMs := duration.Milliseconds()

	if err != nil {
		fmt.Printf("Could not evaluate expression: %v\n", err)
		svc.calculator.SetHasAns(false)
	} else {
		ans = util.CleanFloat(ans)
		fmt.Printf("Result: %g\n", ans)
		svc.calculator.SetHasAns(true)
		svc.calculator.SetAns(ans)
	}
	return ans, err, durationMs
}

func (svc *CalculatorService) SolveEquation(opt int, nums []float64) ([]float64, error, int64) {
	start := time.Now()
	var ans []float64
	var err error
	if opt == 1 {
		ans, err = svc.calculator.SolveLinear(nums)
	} else {
		ans, err = svc.calculator.SolveQuadratic(nums)
	}

	duration := time.Since(start)
	durationMs := duration.Milliseconds()

	if err != nil {
		fmt.Printf("Failed to solve the equation: %v\n", err)
		svc.calculator.SetHasAns(false)
	} else {
		util.PrintSolutions(ans)
		svc.calculator.SetHasAns(true)
	}

	return ans, err, durationMs
}

func (svc *CalculatorService) SolveLinearSystem(matrix [][]float64) ([]float64, error, int64) {
	start := time.Now()
	ans, err := svc.calculator.SolveLinearSystem(matrix)
	duration := time.Since(start)
	durationMs := duration.Milliseconds()
	if err != nil {
		fmt.Printf("Failed to solve the system: %v\n", err)
		fmt.Println("The system may have no solution or infinitely many solutions.")
		svc.calculator.SetHasAns(false)
	} else {
		util.PrintSolutions(ans)
		svc.calculator.SetHasAns(true)
	}

	return ans, err, durationMs
}

func (svc *CalculatorService) SaveAnsToVar(ans []float64, reader *bufio.Reader) bool {
	if !svc.calculator.HasAns() {
		fmt.Println("There is no valid result to save. Please compute a valid result first.")
		return false
	}

	if len(ans) == 0 {
		fmt.Println("There is no result available to save.")
		return false
	}

	chosenAns := ans[0]
	if len(ans) > 1 {
		fmt.Println("\nMultiple results detected.")
		fmt.Println("Select which result to save (1-based index).")
		for {
			index := util.ReadInt(reader, "> ")
			if index <= 0 || index > len(ans) {
				fmt.Printf("Invalid index. Please choose a number between 1 and %d.\n", len(ans))
				continue
			}
			chosenAns = ans[index-1]
			break
		}
	}

	for {
		name := util.ReadLine(reader, "Enter variable name (only letters, no spaces):\n> ")
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Variable name cannot be empty.")
			continue
		}
		if strings.ContainsAny(name, " \t\n") {
			fmt.Println("Variable name cannot contain spaces. Please try again.")
			continue
		}

		if err := svc.calculator.SaveVar(name, chosenAns); err != nil {
			fmt.Printf("Could not save variable: %v\n", err)
			continue
		}

		fmt.Printf("Saved '%s' = %g\n", name, chosenAns)
		return true
	}
}
