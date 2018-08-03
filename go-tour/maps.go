package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

var m1 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex2)

	m["bell labs"] = Vertex2{
		20.23, 1.23,
	}

	fmt.Println(m["bell labs"])

	fmt.Println(m1)

	m2 := make(map[string]int)
	m2["answer"] = 42
	fmt.Println("the value", m2["answer"])

	m2["answer2"] = 43
	fmt.Println(m2)

	delete(m2, "answer")
	fmt.Println(m2)

	v, flag := m2["answer2"]
	fmt.Println("the value:", v, "Present?", flag)

}
