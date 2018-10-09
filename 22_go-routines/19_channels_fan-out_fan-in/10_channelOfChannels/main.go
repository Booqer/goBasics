package main

import (
	"log"
	"sync"
	"time"
)

// the function to be run inside a goroutine. It receives a channel on ch, sleeps for t, then sends t
// on the channel it received
func doStuff(t time.Duration, ch <-chan chan time.Duration) {
	// pulls a value out of the ch argument and assigns it to ac
	// making ac a chan time.Duration
	ac := <-ch
	time.Sleep(t)
	// push a time.Duration into the ac channel
	ac <- t
}

func main() {
	// create the channel-over-channel type
	sendCh := make(chan chan time.Duration)

	recvCh := make(chan time.Duration)

	// use this to block until all goroutines have received the ack and logged
	var wg sync.WaitGroup

	// start up 10 doStuff goroutines
	for i := 0; i < 10; i++ {
		go doStuff(time.Duration(i+1)*time.Second, sendCh)
		// send channels to each doStuff goroutine. doStuff will "ack" by sending its sleep time back
		sendCh <- recvCh
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("slept for %s", <-recvCh)
		}()
	}

	wg.Wait()
}

// From https://www.goin5minutes.com/blog/channel_over_channel/
