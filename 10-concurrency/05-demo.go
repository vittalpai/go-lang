package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	count := make(chan int)
	wg.Add(1)
	go fn(count)
	fmt.Println("Writing the count")
	time.Sleep(4 * time.Second)
	count <- 10
	wg.Wait()
	fmt.Println()
}

func fn(count chan int) {
	fmt.Println("Attempting to read the count")
	max := <-count
	fmt.Println("Just read the data from count channel")
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
