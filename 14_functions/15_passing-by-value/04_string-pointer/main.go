package main

import "fmt"

func main() {
	name := "Derek"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name)
}

func changeMe(g *string) {
	fmt.Println(g)
	fmt.Println(*g)
	*g = "Arnahld"
	fmt.Println(g)
	fmt.Println(*g)
}
