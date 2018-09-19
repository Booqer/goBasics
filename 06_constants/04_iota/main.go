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

const (
	// D is dog, lost in the fog
	D = iota // 0 (notice how it resets with a new declaration)
	// E is every frikkin day
	E // don't need to add the = iota for subsequent declarations
	// F is fog, where dog got lost
	F
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
