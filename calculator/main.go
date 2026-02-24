package main

import (
	"calculator/processor"
	"fmt"
)

func main() {
	var input string
	fmt.Println("======= CLI CALCULATOR ========")

	for {
		fmt.Print("\nPlease input the expression: ")
		fmt.Scanln(&input)

		ans, err := processor.Handle(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Answer: %g\n", ans)
		}

		fmt.Println("\nChoose what you want to do next:")
		fmt.Println("1. Continue")
		fmt.Println("2. Exit")

		var opt int
		_, scanErr := fmt.Scanf("%d\n", &opt)

		if scanErr != nil || opt != 1 {
			if opt == 2 {
				fmt.Println("Goodbye!")
			} else {
				fmt.Println("Not a valid option or Exit, closing...")
			}
			return
		}
	}
}
