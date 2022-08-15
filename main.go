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
	fmt.Println("a + b = ", sum)
	fmt.Println("a - b = ", sub)

}
