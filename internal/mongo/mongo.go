package mongo

import (
	"aandm_server/internal/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func BootstrapDatabase() {

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Config.MongoUser, config.Config.MongoPassword, config.Config.MongoHost, config.Config.MongoPort)))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// @Summary Get all data
// @Description Fetch all data
// @Tags energydata
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/energy-data [get]
func GetData(c *gin.Context) {
	collection := client.Database("your-database-name").Collection("your-collection-name")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// @Summary Get data by id
// @Description Fetch data by id
// @Tags energydata
// @Accept json
// @Produce json
// @Param id path int true "entryId"
// @Success 200 {object} map[string]string
// @Router /api/v1/energy-data:id [get]
func GetDataById(c *gin.Context) {
	id := c.Param("id")
	collection := client.Database("your-database-name").Collection("your-collection-name")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var result bson.M
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
