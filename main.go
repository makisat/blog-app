package main

import (
	// "context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {

	router := gin.Default()

	dotenvInit()

	router.GET("/users", getUsers)
	router.POST("/create-user", addUser)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	mongoInit()

	defer func() {
		if err := clientReference.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	} ()

	fillUsers()
	
	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	mongoInit()

	defer func() {
		if err := clientReference.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	} ()

	collection.InsertOne(ctx, newUser)

	fillUsers()

	c.IndentedJSON(http.StatusCreated, users)
}
