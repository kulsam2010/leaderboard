package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Message struct {
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
	Points   int    `json:"points"`
}

func main() {
	fmt.Println("Consumer app")

	rdb := NewRedisClient()
	ctx := context.Background()

	conn, err := NewRabbitMQConn()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received message %s \n", string(d.Body))
			var msg Message
			err := json.Unmarshal(d.Body, &msg)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Received name= %s , id=%d, points =%d \n", msg.UserName, msg.UserId, msg.Points)
				ProcessMessage(ctx, rdb, msg)
			}

		}
	}()

	fmt.Println("Successfully connected to Rabbit MQ!!")
	fmt.Println("[@] - waiting for messages")
	<-forever
}
