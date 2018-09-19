package main

import "fmt"

func zero(y *int) {
	*y = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
