package main

import "fmt"

func main() {
	q := []int{2, 4, 5, 232, 1, 3}
	fmt.Println(q)

	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
	}

	fmt.Println(s)
}
