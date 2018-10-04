package main

import (
	"fmt"
	"time"
)

func main() {
	cityChan := make(chan string)
	go sendData(cityChan)
	go getData(cityChan)
	time.Sleep(1e9)
}

func sendData(someChan chan string) {
	someChan <- "Washington"
	someChan <- "Tripoli"
	someChan <- "London"
	someChan <- "Beijing"
	someChan <- "Tokyo"
}

func getData(someChan chan string) {
	var input string
	for {
		input = <-someChan
		fmt.Printf("%s ", input)
	}
}
