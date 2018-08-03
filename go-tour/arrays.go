package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}
