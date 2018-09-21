package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9, 11}

	for i, value := range xs {
		fmt.Println(i, " - ", value)
	}
}

// 0  -  1
// 1  -  3
// 2  -  5
// 3  -  7
// 4  -  9
// 5  -  11
