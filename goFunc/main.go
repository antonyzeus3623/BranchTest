package main

import (
	"fmt"
	"time"
)

// 相对应for 循环 goroutine跑到慢 所以这里很大概率只会打印最后一条数据
func goRun1() {
	values := []int{1, 2, 3}
	for _, v := range values {
		go func() {
			fmt.Println("goRun1: ", v)
		}()
	}
}

// 解决1 将参数传入匿名函数 这样参数就可以压栈了
func goRun2() {
	values := []int{1, 2, 3, 4}
	for _, v := range values {
		go func(i int) {
			fmt.Println("goRun2: ", i)
		}(v)
	}
}

// 解决2 将参数赋值给临时变量 也可以将参数压栈
func goRun3() {
	values := []int{1, 2, 3, 4, 5, 6}
	for _, v := range values {
		tmp := v
		go func() {
			fmt.Println("goRun3: ", tmp)
		}()
	}
}

func main() {
	goRun1()
	time.Sleep(time.Second)
	goRun2()
	time.Sleep(time.Second)
	goRun3()
	time.Sleep(time.Second)
}
