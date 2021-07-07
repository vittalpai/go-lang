package main

import "fmt"

func main() {
	/* if. switch, for */
	/* if else contstruct */
	if no := 5; no%2 == 0 {
		fmt.Printf("%d is an even no\n", no)
	} else {
		fmt.Printf("%d is an odd no\n", no)
	}

	//for construct
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//imitating the while construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println(numSum)

	//indefnite loop
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
	}
	fmt.Println(x)

	//switch case construct
	newNo := 7
	switch newNo % 2 {
	case 0:
		fmt.Printf("%d is an even no\n", newNo)
	case 1:
		fmt.Printf("%d is an odd no\n", newNo)
	}

	score := 8
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown Score")
	}

	switch {
	case score >= 0 && score <= 3:
		fmt.Println("Terrible")
	case score >= 4 && score <= 7:
		fmt.Println("Not bad")
	case score >= 8 && score <= 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown Score")
	}

	n := 4
	switch n {
	case 0:
		fmt.Println("terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad")
		fallthrough
	case 8, 9:
		fmt.Println("Good")
		fallthrough
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown Score")
	}

}
