package main

// void Hello(char* s);
import "C"

func main() {
	// C.Hello(C.CString("Hello World!\n"))
	C.Hello(C.CString("Hello World How are you?"))
}
