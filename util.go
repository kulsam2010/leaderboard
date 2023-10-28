package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func Deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)

	err := decoder.Decode(&msg)
	return msg, err

}

func NewRedisClient() *redis.Client {
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	pwd := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	addr := fmt.Sprintf("%s:%d", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
}

func setupRabbitMQChannel(conn *amqp.Connection) (*amqp.Channel, error) {

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error while creating channel")
		fmt.Println(err)
		return nil, err
	}

	queueName := viper.GetString("rabbitmq.queue_name")

	fmt.Println("Queue does not exist. Creating new queue")
	queue, e := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if e != nil {
		fmt.Println("Error while creating new queue")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(queue)
	return ch, nil
}

func NewRabbitMQConn() (conn *amqp.Connection, err error) {
	host := viper.GetString("rabbitmq.host")
	port := viper.GetInt("rabbitmq.port")
	username := viper.GetString("rabbitmq.username")
	pwd := viper.GetString("rabbitmq.pwd")
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, pwd, host, port)
	if addr != "amqp://guest:guest@localhost:5672/" {
		fmt.Println(addr)
		panic("The connections string is incorrect")
	}
	return amqp.Dial(addr)

}

func UserScore(ctx context.Context, rdb *redis.Client, key string, member string) (float64, error) {
	return rdb.ZScore(ctx, key, member).Result()
}

func UpdateScore(ctx context.Context, rdb *redis.Client, key string, member string, points int64) float64 {
	return rdb.ZIncrBy(ctx, key, float64(points), member).Val()
}

func UserRank(ctx context.Context, rdb *redis.Client, key string, member string) (int64, error) {
	rank, err := rdb.ZRevRank(ctx, key, member).Result()
	return rank + 1, err
}
