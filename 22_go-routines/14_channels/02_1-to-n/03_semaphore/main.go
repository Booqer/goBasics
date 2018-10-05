package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		// blocks here until it can take a value off the done channel
		<-done
		// same again
		<-done
		// after taking a value off of done twice, close channel c
		close(c)
	}()

	// the range intrinsically blocks until the channel is closed
	// so we won't reach the end of main() until the channel closes
	for n := range c {
		fmt.Println(n)
	}
}
