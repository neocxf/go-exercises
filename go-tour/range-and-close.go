package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = x+y, x
	}

	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(10, c)

	for i := range c {
		fmt.Println(i)
	}

	v, ok := <-c
	if ok {
		fmt.Println("there are some value", v)
	} else {
		fmt.Println("there are no value")
	}
}
