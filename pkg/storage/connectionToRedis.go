package storage

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectToRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	err:= rdb.Set(ctx, "FIO", "Vladimir Lepeshko", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "FIO").Result()
	if err != nil {
        panic(err)
    }
    fmt.Println(val)

}

