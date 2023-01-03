package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}
func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(string)
		fmt.Println("msg: ", m)
	}
}
func (p *Project) run(msgchan chan interface{}) {
	defer p.deferError()
	for {
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}
func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100)
}
func main() {
	p := new(Project)
	p.Main()
}
