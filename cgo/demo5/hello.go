package main

import "C"

//export hello
func hello(value string) *C.char { // 如果函数有返回值，则要将返回值转换为C语言对应的类型
	return C.CString("hello" + value)
}
func main() {
	// 此处一定要有main函数，有main函数才能让cgo编译器去把包编译成C的库
}
