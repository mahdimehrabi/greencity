package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func ValidateMessageRedis(message []byte) error {
	msg := string(message)
	_, err := rdb.Get(context.Background(), msg).Result()
	if err == nil {
		fmt.Println("validated:", msg)
	}
	return err
}
