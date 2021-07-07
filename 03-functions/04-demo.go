package main

import "fmt"

func main() {
	logger(add, 100, 200)
	logger(subtract, 500, 200)
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
