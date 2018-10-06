package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// randInt := rand.Intn(99)
	factorials := make(chan int)
	semaphore := make(chan bool)
	iterations := 100
	counter := 1
	for i := iterations; i > 0; i-- {
		go func() {
			total := 1
			factorialOf := rand.Intn(9) + 2
			for x := factorialOf; x > 0; x-- {
				total *= x
				// fmt.Println(x, total)
			}
			factorials <- total
			semaphore <- true
		}()
	}
	go func() {
		for i := iterations; i > 0; i-- {
			<-semaphore
		}
		close(factorials)
	}()

	for n := range factorials {
		fmt.Println("Factorial#: ", counter, " : ", n)
		counter++
		// fmt.Println("some number:", rand.Intn(99)+1)
	}
}
