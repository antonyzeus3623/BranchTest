package main

import (
	"BranchTest/RabbitMQ/Simple/Rabbit"
	"fmt"
)

func main() {
	rabbitmq := Rabbit.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.PublishSimple("Hello Wangjie0319!")
	fmt.Println("发送成功！")
}
