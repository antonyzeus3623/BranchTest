package main

//void SayHelloYa(_GoString_ s);
import "C"
import "fmt"

func main() {
	C.SayHelloYa("Hello,World!\n")
}

// export SayHelloYa
func SayHelloYa(s string) {
	fmt.Print(s)
}
