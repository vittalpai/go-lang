package main

import "fmt"

func main() {
	// function assigned to a variable
	var addFn func(int, int) int
	addFn = func(x, y int) int {
		return x + y
	}

	fmt.Println(addFn(100, 200))

	// anonymous functions
	func() {
		fmt.Println("Anonymous fn invoked")
	}()

	func(x, y int) {
		fmt.Println("Subtracting y from x = ", y-x)
	}(100, 200)

	// functions as arguments
	compute(add, 100, 300)
	compute(subtract, 500, 12)

	adder := getAdder()
	fmt.Println(adder(50, 20))

	adderfor100 := getAdderFor(100)
	fmt.Println(adderfor100(200))

}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func compute(fn func(int, int) int, x, y int) {
	fmt.Printf("Computing %d and %d, result = %d\n", x, y, fn(x, y))
}
