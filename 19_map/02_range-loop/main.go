package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Super bad ham!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}

// 1  -  Bonjour!
// 2  -  Buenos dias!
// 3  -  Super bad ham!
// 0  -  Good Morning!
