package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // a's memory address

	var b *int = &a
	fmt.Println(b)  // a's memory address
	fmt.Println(*b) // the value stored at a's memory address

	*b = 42        // b says, "Change the value at this address to 42"
	fmt.Println(a) // 42

	/* This is useful
	we can pass a memory address instead of a bunch of values (we can pass a reference)
	and then we can still change the value of whatever is stored at that memory address
	this makes our programs more performant
	we don't have to pass around large amounts of data
	we only have to pass around addresses

	Everything is PASS BY VALUE in Go, btw
	when we pass a memory address, we are passing a value */
}
