package main

import (
	"fmt"
	"github.com/neocxf/go-exercises/test/iface"
	"github.com/neocxf/go-exercises/test/impl"
	"time"
)

func main() {

	var shape iface.Shaper

	shape = &impl.Circle{R: 3}

	fmt.Println(shape.Area())

	shape = &impl.Rectangle{Width: 2, Length: 3}
	fmt.Println(shape.Area())

	repository := &impl.Repository{}

	retrier := &impl.Retrier{
		RetryCount:   5,
		WaitInterval: time.Second,
		Fetcher:      repository,
	}

	data, err := repository.Fetch(iface.Args{"id": "1"})
	fmt.Printf("#1 repository.Fetch: %v\n", data)

	data, err = retrier.Fetch(iface.Args{})
	fmt.Printf("#2 repository.Fetch: %v\n", data)

	data, err = retrier.Fetch(iface.Args{"id": "1"})
	fmt.Printf("#3 repository.Fetch: %v\n", data)

	if err != nil {
		fmt.Errorf("err is %s\n", err.Error())
	}

}
