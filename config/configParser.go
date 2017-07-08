package config

import (
	"encoding/json"
	"os"
)

func Parse() (conf Config) {
	conf = Config{
		Listen: ListenConfig{
			Listen: ":8081",
		},
		Redis: RedisConfig{
			Address:  "localhost:6379",
			Password: "",
			Database: 0,
		},
		Duration: DurationConfig{
			Auth: "168h",
		},
		PGsql: PGsqlConfig{
			Url: "postgres://dos:dos@localhost:5432/dos?sslmode=disable",
		},
	}

	file, err := os.Open("config.json")
	if err != nil && os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return
}
