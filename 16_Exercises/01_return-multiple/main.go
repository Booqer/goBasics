package main

import "fmt"

func returnMultiple(i int) (int, bool) {
	var isEven bool
	if i%2 == 0 {
		isEven = true
	}
	return i / 2, isEven
}

func main() {

	for x := 1; x < 100; x++ {
		fmt.Println(returnMultiple(x))
	}
}
