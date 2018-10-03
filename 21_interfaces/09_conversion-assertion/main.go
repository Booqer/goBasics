package main

import "fmt"

func main() {
	rem := 5.11
	// rem is of type float64
	fmt.Printf("%T \n", rem)
	// we can convert it to int likes so:
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 9
	// it will show the type of val to be int, but it has an underlying type of interface
	fmt.Printf("%T \n", val)
	// if we try to conver it directly there will be an error:
	// fmt.Printf("%T \n", int(val))
	// so we must instead assert:
	fmt.Printf("%T \n", val.(int))
}
