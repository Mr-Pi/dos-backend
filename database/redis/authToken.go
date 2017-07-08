package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func AuthTokenToUser(token string) (user string) {
	user, err := client.Get("authToken." + token).Result()
	if err != nil && err != redis.Nil {
		fmt.Printf("%x\n", err)
		return
	}
	if err == nil {
		// New TTL
		client.Set("authToken." + token, user, authTokenTTL)
	}
	return
}
