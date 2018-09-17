package main

import "fmt"

func main() {
	// Up here is okay
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	// Down here would fail though
	// x := 42
}

// Package level scope works differently, y is fine down here
var y = 42
