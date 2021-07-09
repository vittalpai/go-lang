package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	var sum int
	ch := make(chan int)
	//ch1 := make(chan int)
	wg.Add(2)

	go add(10, 20, ch)
	go add(100, 200, ch)
	sum += <-ch
	sum += <-ch
	wg.Wait()
	fmt.Println("Sum : ", sum)
}

func add(x int, y int, sum chan int) {
	sum <- x + y
	wg.Done()
}
