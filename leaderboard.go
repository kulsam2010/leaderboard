package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ProcessMessage(ctx context.Context, rdb *redis.Client, msg Message) {
	leaderboardKey := "leaderboard"

	score, err := UserScore(ctx, rdb, leaderboardKey, msg.UserName)

	if err == redis.Nil {
		fmt.Printf("Adding member %s to the leaderboard.\n", msg.UserName)
	} else if err != nil {
		fmt.Printf("Error fetching member value: %v\n", err)
		return
	} else {
		fmt.Printf("Prev. Score of Member %s: %f\n", msg.UserName, score)
	}

	score = score + float64(msg.Points)

	err = rdb.ZAdd(ctx, leaderboardKey, &redis.Z{
		Score:  score,
		Member: msg.UserName,
	}).Err()

	if err != nil {
		fmt.Println("Error Adding score ", err)
	}

	score, _ = UserScore(ctx, rdb, leaderboardKey, msg.UserName)
	rank, _ := UserRank(ctx, rdb, leaderboardKey, msg.UserName)
	fmt.Printf("Score of member %s is %f and rank is %d \n", msg.UserName, score, rank)
}
