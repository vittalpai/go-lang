package main

import (
	"calculator-app/calculator"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	for {
		var input = getUserChoice()
		if input == 5 {
			break
		}
		var n1, n2 = getNumberInputs()
		processInputs(input, n1, n2)
	}
}

func processInputs(input, n1, n2 int) {
	switch input {
	case 1:
		color.Green(fmt.Sprintf("Sum : %d", calculator.Add(n1, n2)))
	case 2:
		color.Green(fmt.Sprintf("Difference : %d", calculator.Subtract(n1, n2)))
	case 3:
		color.Green(fmt.Sprintf("Product : %d", calculator.Multiply(n1, n2)))
	case 4:
		color.Green(fmt.Sprintf("Quotient : %d", calculator.Divide(n1, n2)))
	}
}

func getUserChoice() int {
	input := 0
	color.Red("\n\n\nEnter your choice:\nAdd: 1\nSubtract: 2\nMultiply: 3\nDivide: 4\nExit: 5")
	fmt.Scanf("%d", &input)
	return input
}

func getNumberInputs() (n1, n2 int) {
	color.Yellow("\n\n\nEnter the two numbers [seperated by space]")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return
}
