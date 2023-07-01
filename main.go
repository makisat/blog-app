package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	router := gin.Default()

	dotenvInit()

	router.GET("/users", getUsers)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	mongoInit()

	defer func() {
		if err := clientReference.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	} ()

	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		log.Fatal(err)
	}
	
	c.IndentedJSON(http.StatusOK, users)
}

func addUsers(c *gin.Context) {
	
}
