package main

import "fmt"

func main() {

	a := 43
	b := 43
	c := "Fangly Watson"

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Println("b - ", b)
	fmt.Println("b's memory address - ", &b)
	fmt.Println("c - ", c)
	fmt.Println(c, "'s memory address - ", &c)
}
