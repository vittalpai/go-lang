package main

import (
	"errors"
	"fmt"
)

func main() {
	var n1, n2 int
	/*
		fmt.Println("Enter the number 1")
		fmt.Scanf("%d", &n1)
		fmt.Println("Enter the number 1")
		fmt.Scanf("%d", &n2)
	*/
	fmt.Println("Enter the numbers [seperated by space]")
	fmt.Scanf("%d %d", &n1, &n2)
	result, err := divide(n1, n2)
	if err == nil {
		fmt.Println("Result = ", result)
		return
	}
	fmt.Println(err)
}

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = errors.New("Invalid Argumentss")
		return
	}
	result = x / y
	return
}
