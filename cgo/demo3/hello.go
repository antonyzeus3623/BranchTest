package main

import "C"
import "fmt"

// export Hello
func Hello(s *C.char) {
	fmt.Println(C.GoString(s))
}
