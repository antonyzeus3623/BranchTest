package main

/*
struct Test {
    int a;
    float b;
    double type;
    int size:10;
    int arr1[10];
    int arr2[];
};
int Test_arr2_helper(struct Test * tm ,int pos){
    return tm->arr2[pos];
}
#pragma  pack(1)
struct Test2 {
    float a;
    char b;
    int c;
};
*/
import "C"
import "fmt"

func main() {
	test := C.struct_Test{}
	fmt.Println(test.a)
	fmt.Println(test.b)
	fmt.Println(test._type)
	//fmt.Println(test.size)        // 位数据
	fmt.Println(test.arr1[0])
	//fmt.Println(test.arr)         // 零长数组无法直接访问
	//Test_arr2_helper(&test, 1)

	test2 := C.struct_Test2{}
	fmt.Println(test2.a)
	//fmt.Println(test2.c)          // 由于内存对齐，该结构体部分字段Go无法访问
}
