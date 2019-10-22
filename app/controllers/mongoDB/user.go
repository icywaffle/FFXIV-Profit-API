package mongoDB

import (
	"context"
	"marketboard-backend/app/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserItemStorage struct {
	UserID  string
	Recipes map[string]*models.UserPrice // Stores Recipe, with all related prices
}

// Given a UserStorage collection, it finds all the user's saved item prices
func FindUserItemStorage(userStorage *mongo.Collection, userID string) *UserItemStorage {
	filter := bson.M{"userid": userID}
	var result UserItemStorage
	err := userStorage.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// log.Println(err)
		return nil
	}
	return &result
}

// If a user does not exist in the database, then we would need to give them a table
func InsertNewUserItemStorage(userStorage *mongo.Collection, userPrice *models.UserPrice, userID string) {
	newUserStorage := UserItemStorage{
		UserID:  userID,
		Recipes: make(map[string]*models.UserPrice),
	}
	newUserStorage.Recipes[strconv.Itoa(userPrice.RecipeID)] = userPrice
	userStorage.InsertOne(context.TODO(), newUserStorage)
}

// Once we find a specific user's storage, we just add to it and update it.
func AddUserItem(userStorage *mongo.Collection, userItemStorage *UserItemStorage, userID string, userPrice *models.UserPrice) {
	userItemStorage.Recipes[strconv.Itoa(userPrice.RecipeID)] = userPrice
	filter := bson.M{"userid": userID}
	userStorage.UpdateOne(context.TODO(), filter, bson.D{
		{Key: "$set", Value: userItemStorage},
	})
}
