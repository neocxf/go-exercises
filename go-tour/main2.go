package main

import (
	"fmt"
	"sync"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

const metersToYards float64 = 1.09361

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\t", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\t", TB)

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
