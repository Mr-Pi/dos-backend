package redis

import (
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client
var authTokenTTL time.Duration

func Connect(cfg config.Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	// Parse duration
	authTokenTTL, err = time.ParseDuration(cfg.Duration.Auth)
	if err != nil {
		panic(err)
	}
}
