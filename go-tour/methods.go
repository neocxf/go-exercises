package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// a method is a function with a special receiver argument
// The receiver appears in its own argument list between the func keyword and the method name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{X: 2}

	fmt.Println(v.Abs())

}
