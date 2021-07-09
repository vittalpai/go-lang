/*
modify the program to print the following
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go print("Hello", 1*time.Second, ch1, ch2)
	go print("World", 3*time.Second, ch2, ch1)
	ch1 <- 2
	var input string
	fmt.Scanln(&input)
}

func print(s string, delay time.Duration, ch1 chan int, ch2 chan int) {

	for i := 0; i < 5; i++ {
		<-ch1
		println(s)
		time.Sleep(delay)
		ch2 <- i
	}

}
