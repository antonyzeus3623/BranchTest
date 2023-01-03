package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 以当前时间作为随机种子，生成随机数

	systemByte := r.Uint32()

	log.Println("systemByte: ", systemByte)
}
