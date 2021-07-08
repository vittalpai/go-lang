package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong")
			fmt.Println(r)
		} else {
			fmt.Println("All is well")
		}
	}()

	defer fmt.Println(divide(100, 0))

}

/*
	func divide(x, y int) int {
		return x / y
	}
*/

func divide(x, y int) int {
	if y == 0 {
		panic("Invalid Arguments")
	}
	return x / y
}
