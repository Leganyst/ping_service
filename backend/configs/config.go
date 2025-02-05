package configs

import (
	"os"
)

type ConfigDB struct {
	User     string
	Password string
	DBName   string
	Address  string
	Port     string
}

type Config struct {
	RabbitMQURL string
}

func LoadConfig() *Config {
	return &Config{
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
	}
}

func LoadConfigDB() *ConfigDB {
	// Загружаем переменные окружения из файла .env

	return &ConfigDB{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
