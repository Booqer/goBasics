package main

import "fmt"

func main() {
	// greeting is a func expression
	// ie, an anonymous function assigned to a variable
	// this is the only way to define a function inside of a function
	greeting := func() {
		fmt.Println("Hello World!")
	}

	greeting()
}
