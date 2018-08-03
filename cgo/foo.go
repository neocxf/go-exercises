package cgo

/*
#include <stdio.h>
extern void ACFunction();
*/
import "C"
import "fmt"

func AGoFunction() {
	fmt.Println("AGoFunction()")
}

func Example2() {
	C.ACFunction()
}
