package config

import (
	"errors"
	"os"
)

type Config struct {
	TelegramToken string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		return nil, errors.New("TELE_TOKEN is not set")
	}
	return &Config{TelegramToken: token}, nil
}
