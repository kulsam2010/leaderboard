package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
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
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func NewRabbitMQConn() (conn *amqp.Connection, err error) {
	return amqp.Dial("amqp://leaderboardapp:leaderboardpwd@localhost:5672/")

}

func UserScore(ctx context.Context, rdb *redis.Client, key string, member string) (float64, error) {
	return rdb.ZScore(ctx, key, member).Result()
}

func UserRank(ctx context.Context, rdb *redis.Client, key string, member string) (int64, error) {
	rank, err := rdb.ZRevRank(ctx, key, member).Result()
	return rank + 1, err
}
