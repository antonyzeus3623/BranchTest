package main

//#include "hello.h"
import "C"

func main() {
	// Go 实现的 C 函数
	C.SayHelloWorld(C.CString("Hello World, 你好 世界！"))
}
