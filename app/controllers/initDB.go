package controllers

import (
	"context"
	"log"
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/keys"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongoDB.Collections
var UserStorage *mongo.Collection

// Initializes DB so it would give the Clients so that we can access the database
func InitDB() {
	clientOptions := options.Client().ApplyURI(keys.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("Marketboard")

	DB = mongoDB.Collections{
		Prices:  database.Collection("Prices"),
		Recipes: database.Collection("Recipes"),
		Profits: database.Collection("Profits"),
	}
	UserStorage = database.Collection("UserStorage")

}
