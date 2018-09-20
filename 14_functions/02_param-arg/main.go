package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

// greet is declared with a parameter
// when calling greet, pass in an argument
func greet(name string) {
	fmt.Println(name)
}
