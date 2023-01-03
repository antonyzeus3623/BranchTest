package main

import "log"

func main() {

	s := "OK1"
	len := len([]byte(s + " "))

	log.Println(len)
}
