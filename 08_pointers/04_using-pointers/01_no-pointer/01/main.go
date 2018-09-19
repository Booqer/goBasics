package main

import "fmt"

func zero(x int) int {
	x = 0
	return x
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
	fmt.Println(zero(x))
}
