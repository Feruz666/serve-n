package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	ServerHost string
}

type Gateway struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout struct {
		// Server is the general server timeout to use
		// for graceful shutdowns
		Server time.Duration `yaml:"server"`

		// Write is the amount of time to wait until an HTTP server
		// write opperation is cancelled
		Write time.Duration `yaml:"write"`

		// Read is the amount of time to wait until an HTTP server
		// read operation is cancelled
		Read time.Duration `yaml:"read"`

		// Read is the amount of time to wait
		// until an IDLE HTTP session is closed
		Idle time.Duration `yaml:"idle"`
	} `yaml:"timeout"`
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
