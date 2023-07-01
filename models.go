package main

import (
	"context"
	"log"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type user struct {
	Username string `bson:"username"`
	Age int `bson:"age"`
}

var users []user

func fillUsers() {
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		log.Fatal(err)
	}
}
