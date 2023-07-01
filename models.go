package main

type user struct {
	Username string `bson:"username"`
	Age int `bson:"age"`
}

var users []user
