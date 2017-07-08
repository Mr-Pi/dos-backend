package main

import (
	"fmt"
	"github.com/Mr-Pi/dos-backend/rest"
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/database/redis"
)

func main() {
	cfg := config.Parse()
	redis.Connect(cfg)

	fmt.Println("Router starting")
	rest.InitRouter(cfg)
}
