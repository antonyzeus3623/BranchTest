package main

import (
	"log"
	"strconv"
)

func main() {
	// ParseBool()
	b1 := "T"
	b2, err := strconv.ParseBool(b1)
	if err != nil {
		log.Println("conver to bool failed, err:", err)
	} else {
		log.Printf("type:%T value:%v\n", b2, b2)
	}

	// ParseInt()

}
