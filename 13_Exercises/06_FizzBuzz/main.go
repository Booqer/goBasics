package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		switch {
		case i < 3:
			fmt.Println(i)
		case i%3 == 0:
			fmt.Print("Fizz")
			if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Print("\n")
				break
			}
		case i%5 == 0 || i == 5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
