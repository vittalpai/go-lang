package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10, 20, 213, 12321, 12321, 12312, 312))
	fmt.Println(add(1, 1, 1, 1, 1, 1))
}

func add(nos ...int) int {
	//nos => slice (collection), nos[0], nos[1] ...., len(nos)
	result := 0
	/*
		for i := 0; i < len(nos); i++ {
			result += nos[i]
		}
	*/
	for _, no := range nos {
		//fmt.Println(idx) //prints index
		result += no
	}
	return result
}
