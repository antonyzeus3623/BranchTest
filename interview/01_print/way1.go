package main

import "fmt"

func PrintByThreeChan() {
	letter := make(chan bool)
	number := make(chan bool)
	done := make(chan bool)

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

	go func() {
		let := 'A'
		for {
			select {
			case <-letter:
				if let >= 'Z' {
					done <- true
					return
				}
				fmt.Print(string(let))
				let++
				fmt.Print(string(let))
				let++
				number <- true
			}
		}
	}()

	number <- true
	for {
		select {
		case <-done:
			return
		}
	}
}
