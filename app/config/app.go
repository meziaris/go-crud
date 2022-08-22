package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func (config *Config) Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func New(filenames ...string) *Config {
	godotenv.Load(filenames...)
	return &Config{}
}
