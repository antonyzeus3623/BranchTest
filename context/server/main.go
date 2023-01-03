package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHander(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHander)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
