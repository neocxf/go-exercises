package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	queue := make(chan string)
	queue <- "one"
	queue <- "two"

	wg.Add(1)

	go test(queue)

	wg.Wait()

	go consume(queue)


	fmt.Println("main func finished")
}


func test(c chan <- string) {
	defer wg.Done()

	for i := 0; i < 10; i ++ {
		c <- fmt.Sprintf("str%d", i)
	}
}

func consume(c <- chan string) {
	for elem := range c {
		fmt.Println(elem)
	}
}