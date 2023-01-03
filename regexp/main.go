package main

import (
	"log"
	"regexp"
)

func main() {
	b, err := regexp.Match(".+.log", []byte("log.log"))
	if err != nil {
		log.Println(err)
	}
	log.Println("b:", b)
}
