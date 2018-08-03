package main

import "fmt"

func main() {
	f()

	fmt.Println("Returned normally after f()")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("calling g.")

	g(0)

	fmt.Println("Returned normally after calling g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("pannicking")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defered in g", i)

	fmt.Println("Printing in g", i)

	g(i + 1)
}
