package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

// type cronDemo struct {
// }

// func (c cronDemo) Run() {
// 	log.Println("正在执行定时任务")
// }

// 返回一个支持至 秒 级别的 cron(精确到秒级)
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
func main() {
	// i := 0
	c := newWithSeconds()
	spec := "*/5 * * * * ?" //每秒运行一次

	// c.AddJob(spec, cronDemo{})
	c.AddFunc(spec, func() {
		Demo()
	})
	c.Start()

	defer c.Stop()

	select {}
}

func Demo() {
	log.Println("正在执行定时任务")
}
