package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	a = iota
	b
	c
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
