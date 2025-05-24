package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env file not load: %v", err)
	}

	AppConfig = &Config{MongoURI: getEnv("MONGO_URI")}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s env variable not defined", key)
	}

	return value
}
