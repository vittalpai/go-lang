package main

import "fmt"

func main() {
	var i int = 100
	// addressOfi := &i
	var iPtr *int = &i
	fmt.Println("Address of i = ", iPtr)

	//dereferencing
	fmt.Println("Value at address is = ", *iPtr)

	fmt.Println("Before icrementing, i = ", i)
	i = increment(i)
	fmt.Println("After icrementing, i = ", i)

	fmt.Println("Before icrementing by address, i = ", i)
	incrementByPtr(iPtr)
	fmt.Println("After icrementing by address, i = ", i)

	x, y := 10, 20
	fmt.Println("Before swapping", x, y)
	swap(&x, &y)
	fmt.Println("After swapping", x, y)

	nos := []int{10, 20, 30}
	addValue(&nos, 40)
	fmt.Println(nos)
}

func increment(no int) int {
	fmt.Println("Address of no = ", &no)
	no = no + 1
	return no
}

func incrementByPtr(no *int) {
	fmt.Println("Address of no = ", no)
	*no = *no + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func addValue(nos *[]int, newNumber int) {
	*nos = append(*nos, newNumber)
}
