package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	moo := filter([]int{1, 2, 3, 4, 1}, func(n int) bool {
		return n > 1
	})
	fmt.Println(moo) // [2 3 4]
}
