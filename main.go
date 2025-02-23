package main

import (
	"aandm_server/internal/api"
	"aandm_server/internal/config"
	"aandm_server/internal/mongo"
)

func main() {
	config.LoadConfig()
	mongo.BootstrapDatabase()
	api.BootstrapApi()
}
