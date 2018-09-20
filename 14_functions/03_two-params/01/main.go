package main

import "fmt"

func main() {
	// Here we are calling greet with two arguments
	greet("Jane", "Doe")
}

// greet is declared with two parameters
func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
