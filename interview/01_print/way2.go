package main

import (
	"fmt"
	"sync"
)

func PrintByTwoChan() {
	letter := make(chan bool)
	number := make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		num := 1
		for {
			select {
			case <-number:
				fmt.Print(num)
				num++
				fmt.Print(num)
				num++
				letter <- true
			}
		}
	}()

	wait.Add(1)

	go func(wait *sync.WaitGroup) {
		let := 'A'
		for {
			select {
			case <-letter:
				if let >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(let))
				let++
				fmt.Print(string(let))
				let++
				number <- true
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
}
