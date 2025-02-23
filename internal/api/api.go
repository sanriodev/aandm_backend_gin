package api

import (
	"aandm_server/internal/config"
	"aandm_server/internal/mongo"

	_ "aandm_server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BootstrapApi() {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	api := server.Group("/api/v1")

	// Swagger documentation endpoint
	api.GET("/doc", func(c *gin.Context) {
		c.Redirect(302, "/api/v1/doc/index.html")
	})

	// Swagger documentation handler
	api.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grouping energy data routes under "/energy-data"
	energyDataController := api.Group("/energy-data")
	{
		energyDataController.GET("/", mongo.GetData)
		energyDataController.GET("/:id", mongo.GetDataById)
	}

	server.Run(":" + config.Config.AppPort)
}
