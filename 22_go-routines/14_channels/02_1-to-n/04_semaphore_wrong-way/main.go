package main

import "fmt"

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

	// we block here until done <- true
	// which we will never get to since the goroutines above can't
	// get to the range loop below to hand off to channel c
	<-done
	<-done
	close(c)

	// to unblock above we need to take values off of chan c here
	// but we never get here, because we are blocked above
	for n := range c {
		fmt.Println(n)
	}
}
