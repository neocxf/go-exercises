package main

import (
	"fmt"
	"math/rand"
	"time"
)

func decorator(f func(s string)) func(string) {
	return func(s string) {
		fmt.Println("[decorator]before")
		f(s)
		fmt.Println("[decorator]after")
	}
}

func timeit(f func(s string)) func(string) {
	return func(s string) {
		start := time.Now()
		rand.Seed(start.UnixNano())
		r := rand.Int63n(10)
		fmt.Println("sleeping ", r)
		time.Sleep(time.Duration(r) * time.Second)
		defer func() {
			fmt.Printf("[timtit] took: %v\n", time.Since(start))
		}()

		f(s)

	}
}

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	printSomething("hello")

	decorator(printSomething)("yo yo")

	timeit(printSomething)("yo")

	decorator(timeit(printSomething))("yo2")

}
