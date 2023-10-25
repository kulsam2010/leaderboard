package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ProcessMessage(ctx context.Context, rdb *redis.Client, msg Message) {
	leaderboardKey := "leaderboard"
	score := UpdateScore(ctx, rdb, leaderboardKey, msg.UserName, int64(msg.Points))
	rank, _ := UserRank(ctx, rdb, leaderboardKey, msg.UserName)
	fmt.Printf("Score of member %s is %f and rank is %d \n", msg.UserName, score, rank)
}
