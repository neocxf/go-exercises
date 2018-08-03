package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p1 = &Vertex{1, 2}
)

func c() (i int) {
	// a deferred function increments the return value i after the surrounding function returns.
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println(Vertex{2, 3})

	v := Vertex{1, 2}
	v.X = 3
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, p1)

	fmt.Println(c())
}
