package main

import "fmt"

const PI = 3.14

func main() {
	const world = "world"
	fmt.Println("hello, world", world)
	fmt.Println("happy", PI, "day")

	const truth = true
	fmt.Println("go rule? ", truth)
}
