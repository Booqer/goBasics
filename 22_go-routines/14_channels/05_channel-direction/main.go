package main

import "fmt"

func main() {
	for n := range puller(incrementor()) {
		fmt.Println(n)
	}
}

// the <- here is an optional operator that specifies the channel direction
// if no direction is given, the channel is bidirectional.
// so in this instance incrementor() is returning a channel that we can only
// receive values from, we can not send values to
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
