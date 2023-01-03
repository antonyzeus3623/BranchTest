package main

import (
	"log"
	"strconv"
)

func main() {
	// Atoi, string to int
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		log.Println("string conver to int failed, err: ", err)
	} else {
		log.Printf("type:%T value:%v\n", i1, i1)
	}

	// Itoa, int to string
	i2 := 200
	s2 := strconv.Itoa(i2)
	log.Printf("type:%T value:%v\n", s2, s2)

}
