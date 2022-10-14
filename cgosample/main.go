package main

import "unsafe"

/*
#include <stdio.h>
#include <stdlib.h>
void hello(char *s) { printf("Hello, %s\n", s); }
*/
import "C"

func main() {
	str := C.CString("仙台")
	C.hello(str)
	C.free(unsafe.Pointer(str))
}
