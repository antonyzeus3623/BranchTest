package main

import "log"

type Param map[string]interface{}
type Show struct {
	Param
}

func main() {
	s := new(Show)
	param := make(Param)
	param["RMB"] = 10000
	s.Param = param
	log.Println(s.Param)

}
