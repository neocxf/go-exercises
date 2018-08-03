package main

import (
	"fmt"
	"time"
)

// we can use channels to synchronize execution across goroutines
func worker(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(time.Second)
	fmt.Print("done")

	// Send a value to notify that we’re done.
	done <- true
}

func main() {
	done := make(chan bool, 1)

	// Start a worker goroutine, giving it the channel to notify on.

	go worker(done)

	//Block until we receive a notification from the worker on the channel.
	<-done
}
