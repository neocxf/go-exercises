package main

import "fmt"

// go's return values may be named. If os, they are treated as variables defined at the top of the function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
