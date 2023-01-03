package main

import "BranchTest/RabbitMQ/Simple/Rabbit"

func main() {
	rabbitmq := Rabbit.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.ConsumeSimple()
}
