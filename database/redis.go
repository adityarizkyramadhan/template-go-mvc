package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func NewRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	redisClient = client
	return redisClient
}
