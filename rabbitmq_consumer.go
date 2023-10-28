package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Message struct {
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
	Points   int    `json:"points"`
}

func main() {
	fmt.Println("Consumer app")
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read configuration:", err)
		panic(err)
	}

	rdb := NewRedisClient()
	ctx := context.Background()

	conn, e := NewRabbitMQConn()
	if e != nil {
		fmt.Printf("Error while setting up rabbitMQ %v", e)
		panic(e)
	}
	defer conn.Close()
	ch, err := setupRabbitMQChannel(conn)

	if err != nil {
		fmt.Printf("Error while setting up rabbitMQ Channel %v", err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		viper.GetString("rabbitmq.queue_name"),
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
