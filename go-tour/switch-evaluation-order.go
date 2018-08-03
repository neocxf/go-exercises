package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("what's saturday?")
	today := time.Now().Weekday()

	fmt.Println(today)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Date())

	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")

	}

}
