// work in progress, doesn't compile

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	iterations := 1000 // how many random factorials to calculate
	workerCount := 10  // how many goroutines to use to calculate
	wg.Add(iterations)
	bigChan := make(chan chan int)
	generator := generating(iterations)
	working(workerCount, generator, bigChan)
	merging(bigChan)
	// merging(working(workerCount, generating(iterations)))
	wg.Wait()
}

func generating(iterations int) chan int {
	rand.Seed(time.Now().UTC().UnixNano())
	out := make(chan int)
	go func() {
		for i := 0; i < iterations; i-- {
			// wg.Add(1)
			out <- rand.Intn(23) + 2
		}
		close(out)
	}()
	return out
}

func working(workerCount int, fromGenerator chan int, bigChan chan chan int) {
	toMerger := make(chan chan int)
	for i2 := 0; i2 < workerCount; i2++ {
		go func(fromGenerator chan int) {
			out := make(chan int)
			for n := range fromGenerator {
				out <- factorial(n)
				toMerger <- out
			}
			close(out)
		}(fromGenerator)
	}
}

func merging(bigChan chan chan int) {
	counter := 0
	go func(bigChan chan chan int) {
		for channel := range bigChan {
			for number := range channel {
				counter++
				fmt.Println(counter, "\t", number)
				wg.Done()
			}
		}
	}(bigChan)
}

func factorial(x int) int {
	total := 1
	for i3 := x; i3 > 0; i3-- {
		total *= i3
	}
	return total
}
