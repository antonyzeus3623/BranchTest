package main

/*
#include <stdio.h>
static void HelloPrint(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	C.HelloPrint(C.CString("hello, cgo; hello,world\n"))
}
