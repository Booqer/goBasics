package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Super bad ham!")
	greeting = append(greeting, "Zao'an")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}

// gidday
// 7
// 10
