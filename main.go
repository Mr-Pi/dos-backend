package main

import (
	"fmt"
	"github.com/Mr-Pi/dos-backend/rest"
	"github.com/Mr-Pi/dos-backend/config"
)

func main() {
	cfg := config.Parse()
	fmt.Println("Hello World")
	rest.InitRouter(cfg)
}
