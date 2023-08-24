package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	ServerHost         string
	SentryDSN          string
	JaegerHost         string
	PrometheusHost     string
	PrivateKeyLocation string
	PublicKeyLocation  string
	AccountServiceHost string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("initConfig godotenv load failed")
	}

	cfg = &Config{
		ServerHost: os.Getenv("SERVER_HOST"),
	}

	return cfg, nil
}
