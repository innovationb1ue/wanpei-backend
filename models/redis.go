package models

import (
	"github.com/go-redis/redis/v9"
	"wanpei-backend/server"
)

// NewRedisConn provide the redis Client for the App
func NewRedisConn(settings *server.Settings) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       0,
	})

	return rdb
}
