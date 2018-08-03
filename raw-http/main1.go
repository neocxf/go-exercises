package main

import "github.com/neocxf/go-exercises/raw-http/brokerA"

var bk1 = brokerA.New()
var bk2 = brokerA.New2()

func init() {
	bk1.Execute()
	bk2.Execute()
}

func main() {
}
