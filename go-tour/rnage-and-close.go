package main

import "fmt"

func fibonacci(n uint64, c chan uint64) {
	x, y := uint64(0), uint64(1)
	for i := uint64(0); i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	c := make(chan uint64, 100)
	go fibonacci(uint64(cap(c)), c)

	for i := range c {
		fmt.Println(i)
	}

}
