package main

import "fmt"

func main() {
	name := "Derek"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)
}

func changeMe(g string) {
	fmt.Println(g)
	g = "Arnold"
	fmt.Println(g)
}
