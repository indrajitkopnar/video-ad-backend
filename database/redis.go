package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // e.g. localhost:6379
		Password: "",                      // no password set
		DB:       0,                       // use default DB
	})
}
