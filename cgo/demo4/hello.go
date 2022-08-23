package main

//#include <hello.h>
import "C"
import "fmt"

//export SayHelloWorld
func SayHelloWorld(str *C.char) {
	fmt.Println(C.GoString(str))
}
