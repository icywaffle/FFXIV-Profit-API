package controllers

import (
	"context"
	"ffxiv-profit-api/app/controllers/mongoDB"
	"ffxiv-profit-api/keys"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds the mongo collection for Recipes in our database.
var DB mongoDB.Collections

// UserStorageCollection holds the User Storage collection from our database.
var UserStorageCollection mongoDB.UserStorageCollection

// APIAnalytics should hold client request info to a specific API endpoint
var APIAnalytics *mongo.Collection

// InitDB initializes DB so it would give the Clients so that we can access the database
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
		Recipes: database.Collection("Recipes"),
	}
	UserStorageCollection = mongoDB.UserStorageCollection{
		Collection: database.Collection("UserStorage"),
	}

	APIAnalytics = database.Collection("APIAnalytics")

}
