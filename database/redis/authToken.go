package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"crypto/rand"
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

func NewAuthToken(user string) (token string) {
	for {
		token = generateUUID()
		res, err := client.SetNX("authToken." + token, user, authTokenTTL).Result()
		if err != nil {
			fmt.Printf("%s\n", err)
			token = ""
			return
		}
		if res {
			return
		}
	}
}

func RemoveAuthToken(token string) (success bool) {
	result,err := client.Del("authToken." + token).Result()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	success = result == 1
	return
}

func generateUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}