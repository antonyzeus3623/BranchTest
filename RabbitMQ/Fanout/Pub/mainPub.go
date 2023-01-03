package main

import (
	"BranchTest/RabbitMQ/Fanout/Rabbit"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := Rabbit.NewRabbitMQPubSub("" +
		"newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}

}
