package main

import (
	"os"
	"time"
)

type Config struct {
	BackendURL   string
	RabbitMQURL  string
	PingInterval time.Duration
}

func LoadConfig() Config {
	return Config{
		BackendURL:   os.Getenv("BACKEND_URL"),
		PingInterval: 10 * time.Second,
		RabbitMQURL:  os.Getenv("RABBITMQ_URL"),
	}

}
