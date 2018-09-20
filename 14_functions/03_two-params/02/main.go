package main

import "fmt"

func main() {
	greet("John", "Doe")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
