package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Hi, what's your name?\n")
	fmt.Scan(&name)
	fmt.Println("I hate to say this,", name, ", but you need a shower...")
}
