package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func ValidateMessageRedis(message []byte) {
	msg := string(message)
	_, err := rdb.Get(context.Background(), msg).Result()
	if err == nil {
		fmt.Println("validated:", msg)
	} else {
		fmt.Println("failed to validate:", msg, err)
	}
}
