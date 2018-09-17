package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
	/* the following will not work since within block level scope order matters
	fmt.Println(z)
	z := 34
	*/
}

func foo() {
	fmt.Println(x)
	// the following will not work since y was initialized in a separate function
	// fmt.Println(y)
}
