package main

import "fmt"

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
		fmt.Println("Sum : ", add(n1, n2))
	case 2:
		fmt.Println("Difference : ", subtract(n1, n2))
	case 3:
		fmt.Println("Product : ", multiply(n1, n2))
	case 4:
		fmt.Println("Quotient : ", divide(n1, n2))
	}
}

func getUserChoice() int {
	input := 0
	fmt.Println("\n\n\nEnter your choice:\nAdd: 1\nSubtract: 2\nMultiply: 3\nDivide: 4\nExit: 5")
	fmt.Scanf("%d", &input)
	return input
}

func getNumberInputs() (n1, n2 int) {
	fmt.Println("\n\n\nEnter the two numbers [seperated by space]")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
