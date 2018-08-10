package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	messages := make(chan int)

	done := make(chan bool)

	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
		done <- true
	}()

	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
		done <- true
	}()

	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
		done <- true
	}()

	go func() {
		for i := range messages {
			fmt.Println(i)

		}
	}()

	for i := 0; i < 3; i++ {
		<-done
	}
}

func Test2(t *testing.T) {
	messages := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}()

	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)

	}
}

func Test3(t *testing.T) {
	urls := []string{
		"http://www.reddit.com/r/aww.json",
		"http://www.reddit.com/r/funny.json",
		"http://www.reddit.com/r/programming.json",
	}

	resc, errc := make(chan string), make(chan error)

	for _, url := range urls {
		go func(url string) {
			body, err := fetch(url)
			if err != nil {
				errc <- err
				return
			}

			resc <- string(body)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case res := <-resc:
			fmt.Println(res)
		case err := <-errc:
			fmt.Println(err)
		}
	}
}

func fetch(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}
