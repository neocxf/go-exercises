package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		time.Sleep(time.Second)
		fmt.Println("wake up")
	}
}
