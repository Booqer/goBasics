package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int)
	go pump(intChan)
	go suck(intChan)
	time.Sleep(1e9)
}

func pump(someChan chan int) {
	for i := 0; ; i++ {
		someChan <- i
	}
}

func suck(someChan chan int) {
	for {
		fmt.Println(<-someChan)
	}
}
