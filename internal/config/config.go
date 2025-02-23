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
	ModbusHost := os.Getenv("MODBUS_HOST")
	ModbusPort := os.Getenv("MODBUS_PORT")

	*Config = types.AppConfig{AppPort: AppPort, MongoHost: MongoHost, MongoPort: MongoPort, MongoUser: MongoUser, MongoPassword: MongoPassword, ModbusHost: ModbusHost, ModbusPort: ModbusPort}
	return *Config
}
