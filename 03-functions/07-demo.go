package main

import "fmt"

func main() {
	/*
		defer func() {
			fmt.Println("deferred from main")
		}()
	*/
	defer fmt.Println("deferred from main")
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer fmt.Println("deferred f11")
	defer fmt.Println("deferred f12")
	defer fmt.Println("deferred f13")
	fmt.Println("f1 invoked")
	msg := f2()
	fmt.Println(msg)
}

func f2() (message string) {
	message = "vittal"
	defer func() {
		message = "vittal deferred"
		fmt.Println("deferred f21")
	}()
	defer fmt.Println("deferred f22")
	defer fmt.Println("deferred f23")
	fmt.Println("f2 invoked")
	return
}
