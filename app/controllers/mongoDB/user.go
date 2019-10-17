package mongoDB

import (
	"context"
	"marketboard-backend/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserItemStorage struct {
	UserID int
	Recipe map[int]*models.UserPrice // Stores Recipe, with all related prices
}

// Given a UserStorage collection, it finds all the user's saved item prices
func FindUserItemStorage(userStorage *mongo.Collection, userID int) *UserItemStorage {
	filter := bson.M{"UserID": userID}
	var result UserItemStorage
	err := userStorage.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// log.Println(err)
		return nil
	}
	return &result
}

// If a user does not exist in the database, then we would need to give them a table
func InsertNewUserItemStorage(userStorage *mongo.Collection, userID int) {
	newUserStorage := UserItemStorage{
		UserID: userID,
		Recipe: make(map[int]*models.UserPrice),
	}
	userStorage.InsertOne(context.TODO(), newUserStorage)
}

// Once we find a specific user's storage, we just add to it and update it.
func AddUserItem(userStorage *mongo.Collection, userItemStorage *UserItemStorage, userID int, userPrice *models.UserPrice) {
	userItemStorage.Recipe[userPrice.RecipeID] = userPrice
	filter := bson.M{"UserID": userID}
	userStorage.UpdateOne(context.TODO(), filter, bson.D{
		{Key: "$set", Value: userItemStorage},
	})
}
