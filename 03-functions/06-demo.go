package main

import "fmt"

func main() {
	increment, decrement := getCounters()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(increment())
	fmt.Println(increment())
}

func getCounters() (func() int, func() int) {
	var count int
	return func() int {
			count++
			return count
		}, func() int {
			count--
			return count
		}
}
