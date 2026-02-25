package main

import (
	"bufio"
	"calculator/internal/utils"
	"calculator/processor"
	"fmt"
	"os"
)

func main() {
	fmt.Println("CLI CALCULATOR")
	fmt.Println("Type a menu number and press Enter.")

	calculator := processor.NewCalculator()
	reader := bufio.NewReader(os.Stdin)

	for {
		opt := utils.ReadInt(reader,
			"\nMain Menu\n"+
				"1. Expression\n"+
				"2. Equation\n"+
				"3. Linear System\n"+
				"4. Exit\n"+
				"> ",
		)

		switch opt {
		case 1:
			runExpression(calculator, reader)
		case 2:
			runEquation(calculator, reader)
		case 3:
			runLinearSystem(calculator, reader)
		case 4:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Printf("Invalid option (%d). Please choose 1, 2, 3, or 4.\n", opt)
		}
	}
}
