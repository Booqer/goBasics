package main

import (
	"fmt"
	"time"
)

func main() {
	chatChan := make(chan string)
	go supBrah(chatChan)
	go heyBro(chatChan)
	time.Sleep(40e9)
	fmt.Println("main() done.")
}

func supBrah(chatChan chan string) {
	moo := "moo"
	fmt.Println("Sup brah, I'm mooing at you.")
	chatChan <- moo
}

func heyBro(chatChan chan string) {
	fmt.Println("Hey bro, I hear you.")
	time.Sleep(15e9)
	moo := <-chatChan
	fmt.Println("Did you drop this? ", moo)
	time.Sleep(15e9)
	fmt.Println("Shoo, be gone with ye")
}
