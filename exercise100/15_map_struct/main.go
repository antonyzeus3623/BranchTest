package main

import "log"

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	for i := range m {
		log.Printf("%#v", m[i])
	}

}
