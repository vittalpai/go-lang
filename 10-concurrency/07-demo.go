package main

import (
	"time"
)

func main() {
	data := []int{2, 1, 2321, 123, 12, 321, 321, 2, 321, 4, 2, 12, 232}
	firstSet := data[:len(data)/2]
}

func sum(resultCh, ch2 chan int) {

	for i := 0; i < 5; i++ {
		<-ch1
		println(s)
		time.Sleep(delay)
		ch2 <- i
	}

}
