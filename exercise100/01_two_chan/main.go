// 使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟
// 效果: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
package main

import (
	"fmt"
	"sync"
)

func main() {
	letter := make(chan bool, 1)
	number := make(chan bool, 1)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
