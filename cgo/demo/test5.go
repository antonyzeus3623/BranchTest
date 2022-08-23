package main

// int SayHello();
import "C"
import "fmt"

func main() {
	// C.puts(C.CString("Hello, Cgo\n"))

	ret := C.SayHello()
	fmt.Println(ret)
}
