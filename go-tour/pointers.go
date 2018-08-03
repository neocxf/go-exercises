package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(p)

	var q *int
	fmt.Println(q)

	if q == nil {
		fmt.Println("q is nil")
	}

	q = &i

	fmt.Printf("%T\n", *q)

	var m = q
	fmt.Println(*m)

}
