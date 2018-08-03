package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (myreader *rot13Reader) Read(p []byte) (n int, err error) {
	n, err0 := myreader.r.Read(p)

	if err0 != nil {
		return n, err0
	}

	for i := 0; i < n; i++ {
		switch {
		case p[i] >= 65 && p[i] <= 77:
			p[i] += 13
		case p[i] >= 97 && p[i] <= 109:
			p[i] += 13
		case 78 <= p[i] && p[i] <= 90:
			p[i] -= 13
		case p[i] >= 110 && p[i] <= 122:
			p[i] -= 13

		}

		//if (65 <= p[i] && p[i] <=77) || (p[i] >=97 && p[i] <=109) {
		//	p[i] = p[i] + 13
		//	continue
		//}
		//
		//if (78 <= p[i] && p[i] <=90) || (p[i] >=110 && p[i] <=122) {
		//	p[i] = p[i] - 13
		//	continue
		//}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
