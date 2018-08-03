package main

import (
	"fmt"
	"github.com/Go-zh/tour/reader"
	"io"
)

type MyReader struct {
}

func (myreader MyReader) Read(p []byte) (n int, err error) {

	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}

	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})

	var myreader io.Reader

	myreader = &MyReader{}

	b := make([]byte, 8)

	for i := 0; i < 3; i++ {
		n, err := myreader.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}

}
