package main

import "fmt"

func zero(y int) {
	fmt.Printf("%p\n", &y) // address in func zero
	fmt.Println(&y)        // address in func zero
	y = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
