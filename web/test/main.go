package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mux sync.RWMutex

func main() {

	queue := make(chan string)
	quit := make(chan struct{})

	wg.Add(1)

	go produce(queue, quit)

	go consume(queue, quit)

	wg.Wait()

	fmt.Println("main func finished")
}

func produce(c chan<- string, quit chan<- struct{}) {

	for i := 0; i < 10; i++ {
		fmt.Println("producing msg: ", fmt.Sprintf("str%d", i))
		c <- fmt.Sprintf("str%d", i)
	}

	quit <- struct{}{}
}

func consume(c <-chan string, quit <-chan struct{}) {
consumeLabel:
	for {
		select {
		case elem := <-c:
			fmt.Println(elem)
		case <-quit:
			fmt.Println("exiting ....")
			break consumeLabel
		}
	}
	wg.Done()
}

func consume2(c <-chan string, quit <-chan struct{}) {
	for {
		select {
		case elem := <-c:
			fmt.Println(elem)
		case <-quit:
			fmt.Println("exiting ....")
			wg.Done()
			return
		}
	}
}
