package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Leaderboard app")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()
	fmt.Println("Successfully connected to Rabbit MQ!!")
	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	message := "Hello RMQ!"
	publishToRmq(ch, message)
}

func publishToRmq(ch *amqp.Channel, message string) bool {
	queue, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(queue)
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Successfully published the message to RMQ")
	return true
}
