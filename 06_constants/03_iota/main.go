package main

import "fmt"

const (
	// A an ardvark, a screaming ardvark
	A = iota // 0
	// B a hefty tub of beer
	B = iota // 1
	// C is chill, does what it will
	C = iota // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
