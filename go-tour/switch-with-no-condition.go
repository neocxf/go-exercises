package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// a clean way to write long if-then-else chains
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}
