package main

/*
int SayHello(char* buff, int len) {
    char hello[] = "Hello Cgo!";
    int movnum = len < sizeof(hello) ? len:sizeof(hello);
    memcpy(buff, hello, movnum);                        // go字符串没有'\0'，所以直接内存拷贝
    return movnum;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 切片
func main() {
	buff := make([]byte, 8)
	C.SayHello((*C.char)(unsafe.Pointer(&buff[0])), C.int(len(buff)))
	a := string(buff)
	fmt.Println(a)
}
