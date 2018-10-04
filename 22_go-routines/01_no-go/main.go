package main

import "fmt"

func main() {
	// with no "go" we have no concurrency built in
	// so foo will run completely then bar
	foo()
	bar()
	// if we were to add the go keyword as below but change nothing else
	// then there would be three seperate routines running at once:
	// foo(), bar(), and main()
	// while foo() and bar() are spooling up, main() will finish executing
	// and the program will be over before foo or bar have a chance to run
	// go foo()
	// go bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}
