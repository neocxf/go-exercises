package main

import (
	"fmt"
	"math/rand"
	"time"
)

type somedata struct {
	s string
}

func (d *somedata) returnStrWithPrefix(prefix string) string {
	return fmt.Sprintf("%s %s", prefix, d.s)
}

type DoPrintSomething func(string) string

func decorate(f DoPrintSomething) DoPrintSomething {
	return func(s string) string {
		start := time.Now()

		defer func() {
			fmt.Println("[decorate] took: ", time.Since(start))
		}()

		rand.Seed(start.UnixNano())

		time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

		return f(s)
	}
}

func main() {

	d := somedata{s: "i'm not decorated\n"}
	ds := d.returnStrWithPrefix(">>")
	fmt.Println(ds)

	e := somedata{s: "i now am decorated... woohoo!"}
	es := decorate(e.returnStrWithPrefix)("==>")
	fmt.Println(es)
}
