package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		panic("could not establish AMQP connection: " + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open channel: " + err.Error())
	}

	// declare queue
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}

	// bind queue
	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		panic("error while binding the queue: " + err.Error())
	}

	// consume
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		panic("error while consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		fmt.Println("message received: ", string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()
}
