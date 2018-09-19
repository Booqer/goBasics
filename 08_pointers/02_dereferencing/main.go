package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // a's memory address

	var b *int = &a // assign b as a's memory adress
	fmt.Println(b)  // a's memory address
	fmt.Println(*b) // 43 - gives the value at that memory address

	a += 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	/* b is an int pointer;
	it points to the memory address where an int is stored
	to see the value in that memory address, add an * in front of b
	this is known as dereferencing
	the * is an operator in this case
	*/
}
