package main

import (
	"fmt"
	"github.com/Mr-Pi/dos-backend/config"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/database/redis"
	"github.com/Mr-Pi/dos-backend/rest"
)

func main() {
	cfg := config.Parse()
	redis.Connect(cfg)
	pgsql.Connect(cfg)

	fmt.Println(cfg)
	fmt.Println("Router starting")
	rest.InitRouter(cfg)
}
