package main

import "fmt"

func variadic(x ...int) int {
	var theBigOne int
	for _, v := range x {
		if v > theBigOne {
			theBigOne = v
		}
	}
	return theBigOne
}

func main() {
	fmt.Println(variadic(111, 1, 11, 54, 73, 88))
}
