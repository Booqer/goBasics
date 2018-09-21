package main

import "fmt"

func main() {
	greeting := []string{
		"Good Morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}

// 0 Good Morning!
// 1 Bonjour!
// 2 dias!
// 3 Bongiorno!
// 4 Ohayo!
// 5 Selamat pagi!
// 6 Gutten morgen!
// Good Morning!
// Bonjour!
// dias!
// Bongiorno!
// Ohayo!
// Selamat pagi!
// Gutten morgen!
