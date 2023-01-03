package main

import "BranchTest/RabbitMQ/Fanout/Rabbit"

func main() {
	rabbitmq := Rabbit.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
