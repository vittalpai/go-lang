package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	//ch1 := make(chan int)
	wg.Add(3)

	go add(10, 20, ch)
	go add(100, 200, ch)
	res1 := <-ch
	res2 := <-ch
	go add(res1, res2, ch)
	fmt.Println("Sum : ", <-ch)
	wg.Wait()
}

func add(x int, y int, sum chan int) {
	sum <- x + y
	wg.Done()
}
