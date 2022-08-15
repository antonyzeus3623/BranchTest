package main

import (
	"fmt"

	"BranchTest/add"
	"BranchTest/sub"
)

func main() {
	a := 5
	b := 2

	sum := add.Add(a, b)
	sub := sub.Sub(a, b)

	fmt.Println("加法计算 ：a + b = ", sum)
	fmt.Println("减法计算：a - b = ", sub)

}
