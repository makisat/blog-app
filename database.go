package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()
var db_string string
var clientReference *mongo.Client

func mongoInit() {

	clientOptions := options.Client().ApplyURI(db_string)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	clientReference = client

	collection = client.Database("user").Collection("users")
}

func dotenvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db_string = os.Getenv("DB_STRING")
}
