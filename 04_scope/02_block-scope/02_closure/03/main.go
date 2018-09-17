package main

import "fmt"

func main() {
	x := 0
	// this is an anonymous function... a function without a name
	// to put a function inside a function you have to do it like this
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable
that variable would need to be package scope
*/
