package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConsumeMessages (channel *amqp.Channel) {

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"test", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func ()  {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}