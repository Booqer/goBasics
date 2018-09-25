package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour",
		2: "Buenos dias!",
		3: "Sawat dee",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}

// map[3:Sawat dee 0:Good morning! 1:Bonjour 2:Buenos dias!]
// val:  Buenos dias!
// exists:  true
// map[0:Good morning! 1:Bonjour 3:Sawat dee]
