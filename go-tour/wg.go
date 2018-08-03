package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func routine(idx int) {
	defer wg.Done()

	fmt.Println("routine finished", idx)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routine(i)

	}

	wg.Wait()

	fmt.Println("main finished")

}
