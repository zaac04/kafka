package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis Cache

func ConnectRedis() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		fmt.Println(err)
	}
	Redis.client = redis.NewClient(&redis.Options{
		Addr:           os.Getenv("REDIS_URL"),
		Password:       os.Getenv("REDIS_PASS"),
		DB:             db,
		PoolSize:       10,
		MaxActiveConns: 10,
	})
}
