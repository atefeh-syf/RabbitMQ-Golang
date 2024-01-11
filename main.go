package main

import (
	"fmt"

	"github.com/atefeh-syf/RabbitMQ-Golang/consumer"
	"github.com/atefeh-syf/RabbitMQ-Golang/producer"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel := producer.PublishingMessage(connection)
	consumer.ConsumeMessages(channel)
}