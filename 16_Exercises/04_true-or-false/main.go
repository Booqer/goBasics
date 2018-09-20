package main

import "fmt"

func main() {
	var gonnaBeTrue bool
	if (true && true) || (false && true) || !(false && false) {
		gonnaBeTrue = true
	}
	fmt.Println(gonnaBeTrue)
}
