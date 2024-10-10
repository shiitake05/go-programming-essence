package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello(char *str) {
	printf("Hello, %s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	s := "world"

	ptr := C.CString(s)
	defer C.free(unsafe.Pointer(ptr))
	C.hello(ptr)
}
