package main

import "fmt"

// test where an interface value holds a specific type,
// a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
