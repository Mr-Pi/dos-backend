package config

type Config struct {
	Listen ListenConfig `json:"listen"`
	Redis RedisConfig `json:"redis"`
	Duration DurationConfig `json:"duration"`
}
