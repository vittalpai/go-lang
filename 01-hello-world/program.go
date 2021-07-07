//package declaration
package main

//import other packages
import "fmt"

//package level variables/types

//package init function
func main() {
	/*
		var message string
		message = "Hello World"
	*/

	/*
		var message string = "Hello World"
	*/

	/*
	 var message = "Hello World"
	*/

	// following syntx works only in a function
	message := "Hello World"
	fmt.Println(message)

	//using multiple variables
	/*
		var (
			name string
			msg  string
		)
	*/

	/*
		var (
			name, msg string
		)
	*/

	/*
		var name,msg string
		name = "Vittal"
		msg = "Welcome to GoLang"
	*/

	/*
		var name, msg string = "Vittal", "Welcome to Golang"
	*/

	/*
		var name, msg = "Vittal", "Welcome to Golang"
	*/

	name, msg := "Vittal", "Welcome to Golang"
	fmt.Println(name, msg)
}

//other functions
