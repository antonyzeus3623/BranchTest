package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case student:
		msg.Name = "lili"
		fmt.Println(msg.Name)
	}
}

func main() {
	x := "student"
	zhoujielun(x)
}
