package main

import "fmt"

type somedata struct {
	s string
}

func (d *somedata) printSomething() {
	fmt.Println(d.s)
}

type DoPrintSomething func()

func decorate(f DoPrintSomething) DoPrintSomething {
	return func() {
		fmt.Println("[decorate] before")
		f()
		fmt.Println("[decorate] after")
	}
}

func main() {

	d := somedata{s: "i am not decorated\n"}

	d.printSomething()

	e := somedata{s: "i now am decorated... woohoo!"}
	decorate(e.printSomething)()
}
