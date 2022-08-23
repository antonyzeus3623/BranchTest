package main

/*
#include <stdio.h>

void printint(int v){
	printf("printint:%d\n",v);
}
*/
import "C"

func main() {
	v := 50
	C.printint(C.int(v))
}
