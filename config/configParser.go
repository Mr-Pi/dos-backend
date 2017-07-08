package config

import (
	"os"
	"encoding/json"
)

func Parse() (conf Config) {
	conf = Config{
		Listen: ListenConfig{
			Listen: ":8081",
		},
	}

	file,err := os.Open("config.json")
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
