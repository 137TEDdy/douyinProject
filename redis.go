package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	var _ *redis.Client

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Network:  "tcp",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		fmt.Println("Redis连接失败:", err)
		return
	}
	err := rdb.Set(ctx, "123", 123, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println(rdb.Get(ctx, "123"))
	fmt.Println("Redis连接成功")
}
