package config

import (
	"aandm_server/internal/types"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config = new(types.AppConfig)

func LoadConfig() types.AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	AppPort := os.Getenv("APP_PORT")
	MongoHost := os.Getenv("MONGO_HOST")
	MongoPort := os.Getenv("MONGO_PORT")
	MongoUser := os.Getenv("MONGO_USERNAME")
	MongoPassword := os.Getenv("MONGO_PASSWORD")
	MongoDatabase := os.Getenv("MONGO_DATABASE")

	*Config = types.AppConfig{AppPort: AppPort, MongoHost: MongoHost, MongoPort: MongoPort, MongoUser: MongoUser, MongoPassword: MongoPassword, MongoDatabase: MongoDatabase}
	return *Config
}
