package main

import "fmt"

func deferReturn() string {
	fmt.Println("deferred call arguments are evaluated first, the return result will be evaluated until the calling defer function returns")
	return "world"
}

func main() {
	defer fmt.Println(deferReturn())

	fmt.Println("hello")
}
