package main

import "fmt"

func main() {
	sum := 8
	for i := 8; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
