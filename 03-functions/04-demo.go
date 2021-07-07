package main

import "fmt"

func main() {
	logger(add, 100, 200)
	logger(subtract, 500, 200)
	loggerAdd := getloggerfun(add, "Add")
	fmt.Println(loggerAdd(10, 15))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func logger(fn func(int, int) int, x, y int) {
	fmt.Println("Begin invocation")
	fmt.Println(fn(x, y))
	fmt.Println("After invocation")
}

func getloggerfun(fn func(int, int) int, operName string) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("Begin invocation")
		result := fn(x, y)
		fmt.Println("Result =", result)
		fmt.Println("After invocation")
		return result
	}
}
