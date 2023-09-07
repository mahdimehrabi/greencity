package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func StoreMessageRedis(message []byte) {
	err := rdb.Set(context.Background(), string(message), "1", 10*time.Minute).Err()
	if err != nil {
		fmt.Println("failed to store message to redis", string(message), err)
	}
}
