package main

import "fmt"

func foo(x ...int) {
	for _, v := range x {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
