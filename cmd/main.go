package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	s, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}
