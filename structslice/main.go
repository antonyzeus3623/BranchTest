package main

import (
	"encoding/json"
	"log"
)

type Student struct {
	Class string     `json:"class"`
	Info  []*StuInfo `json:"studentInfo"`
}

type StuInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	stu := new(Student)
	var stuInfo []*StuInfo
	stuInfo = append(stuInfo, &StuInfo{Name: "小莉", Age: 19, Score: 100})
	stuInfo = append(stuInfo, &StuInfo{Name: "平平", Age: 18, Score: 98})

	stu.Class = "大学一年级1班"
	stu.Info = append(stu.Info, stuInfo...)

	data, _ := json.Marshal(stu)

	log.Println(stu)
	log.Printf("%#v\n", &stu.Info)

	log.Println(string(data))
}
