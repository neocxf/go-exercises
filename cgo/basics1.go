package cgo

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func Example() {
	cs := C.CString("hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
