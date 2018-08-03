package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan string)

	msg := "hi"

	go func() {
		messages <- "hello"
		signals <- "sig"

	}()

	time.Sleep(1 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("receive message1", msg)
	default:
		fmt.Println("no message received")
	}

	go func() {
		<-messages
	}()

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
		//default:
		//	fmt.Println("no message sent")
	}

	select {
	case sig := <-signals:
		fmt.Println("receive signal", sig)
	case msg := <-messages:
		fmt.Println("received message2", msg)
	default:
		fmt.Println("no activity")
	}

}
