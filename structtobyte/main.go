package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type TestStructTobytes struct {
// 	data int64
// }
// type SliceMock struct {
// 	addr uintptr
// 	len  int
// 	cap  int
// }

type MyStruct struct {
	A int
	B int
}

// var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = int(unsafe.Sizeof(MyStruct{}))
	x.Cap = int(unsafe.Sizeof(MyStruct{}))
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func main() {
	test := &MyStruct{
		A: 100,
		B: 50,
	}
	buf := MyStructToBytes(test)
	fmt.Println("byte:", buf)

	// var testStruct = &TestStructTobytes{100}

	// Len := unsafe.Sizeof(*testStruct)
	// testBytes := &SliceMock{
	// 	addr: uintptr(unsafe.Pointer(testStruct)),
	// 	cap:  int(Len),
	// 	len:  int(Len),
	// }
	// data := *(*[]byte)(unsafe.Pointer(testBytes))
	// fmt.Println("[]byte is : ", data)
	// var ptestStruct *TestStructTobytes = *(**TestStructTobytes)(unsafe.Pointer(&data))
	// fmt.Println("ptestStruct.data is : ", ptestStruct.data)

}
