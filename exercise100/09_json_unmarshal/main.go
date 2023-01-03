package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	js := `{
	"Name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("people: % v", p)
}
