package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host string
	Port string
}

type Config struct {
	Server *ServerConfig
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewConfig(config *Config) *Config {
	defaultConfig := &Config{
		Server: &ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "0.0.0.0"),
		},
	}

	if config != nil {
		if config.Server != nil {
			defaultConfig.Server = config.Server
		}
	}

	return &Config{
		Server: defaultConfig.Server,
	}
}
