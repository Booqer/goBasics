package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is the length - number of elements referred to by the slice
	// 5 is the capacity - number of elements in the underlying array

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Super bad ham!")

	fmt.Println(greeting[3])
}
