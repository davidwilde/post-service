package redis_backend

import (
	"gopkg.in/redis.v3"
)

func NewRedisClient() *redis.Client {
	opts := &redis.Options{
		Addr:     "db:6379",
		Password: "",
		DB:       0,
	}
	return redis.NewClient(opts)
}
