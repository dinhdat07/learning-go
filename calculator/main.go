package main

import (
	"bufio"
	"calculator/internal/utils"
	"calculator/processor"
	"fmt"
	"os"
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
			runLinearSystem(calculator, reader)
		case 4:
			return
		default:
			fmt.Println("Not a valid option, please choose 1/2/3.")
		}
	}

}
