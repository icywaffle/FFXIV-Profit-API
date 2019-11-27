package mongoDB

import (
	"context"
	"ffxiv-profit-api/app/models"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserStorageCollection stores a user's storage collection from the database, with a few methods.
type UserStorageCollection struct {
	*mongo.Collection
}

// UserStorageCollectionHandler handles the user storage.
type UserStorageCollectionHandler interface {
	FindUserItemStorage(userID string) *UserItemStorage
	InsertNewUserItemStorage(UserSubmission *models.UserSubmission, userID string)
	AddUserItem(userItemStorage *UserItemStorage, userID string, UserSubmission *models.UserSubmission)
}

// UserItemStorage is an object that holds for a specific UserID, all their posted prices and profits.
type UserItemStorage struct {
	UserID  string
	Prices  map[string]models.UserPrices
	Profits map[string]models.UserProfits
}

// FindUserItemStorage finds all items that are stored in the User's Document
func (coll UserStorageCollection) FindUserItemStorage(userID string) *UserItemStorage {
	filter := bson.M{"userid": userID}
	var result UserItemStorage
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// log.Println(err)
		return nil
	}
	return &result
}

// InsertNewUserItemStorage creates a new UserItem storage if a user does not exist in the database.
func (coll UserStorageCollection) InsertNewUserItemStorage(UserSubmission *models.UserSubmission, userID string) {
	newUserStorage := UserItemStorage{
		UserID:  userID,
		Prices:  make(map[string]models.UserPrices),  // Key: ItemID
		Profits: make(map[string]models.UserProfits), // Key: RecipeID
	}
	AddUserInfoToMap(&newUserStorage, UserSubmission)

	coll.InsertOne(context.TODO(), newUserStorage)
}

// AddUserItem adds to the user storage and update it.
func (coll UserStorageCollection) AddUserItem(userItemStorage *UserItemStorage, userID string, UserSubmission *models.UserSubmission) {
	AddUserInfoToMap(userItemStorage, UserSubmission)
	filter := bson.M{"userid": userID}
	coll.UpdateOne(context.TODO(), filter, bson.D{
		{Key: "$set", Value: userItemStorage},
	})
}

// DeleteUserItem removes an user's item storage permanently.
func (coll UserStorageCollection) DeleteUserItem(userID string) {
	filter := bson.M{"userid": userID}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

// AddUserInfoToMap places all the UserSubmission into a user's storage document.
func AddUserInfoToMap(newUserStorage *UserItemStorage, UserSubmission *models.UserSubmission) {
	// We add the main recipe first, then it's materials.
	// If the material is craftable, then it'll have it's separate call to this function
	mainProfits := models.UserProfits{
		RecipeID:         UserSubmission.RecipeID,
		ItemID:           UserSubmission.ItemID,
		ItemName:         UserSubmission.ItemName,
		IconID:           UserSubmission.IconID,
		MaterialCosts:    UserSubmission.MaterialCosts,
		IngredientItemID: UserSubmission.IngredientItemID,
		Profits:          UserSubmission.Profits,
		ProfitPercentage: UserSubmission.ProfitPercentage,
	}

	// Prices is a bit tricky.
	// Since it has an old map with information, we cannot just overwrite.
	// We have to take the old map if it exits
	newUsedFor := make(map[string]bool)
	ItemID := strconv.Itoa(UserSubmission.ItemID)
	prevPrices, ok := newUserStorage.Prices[ItemID]
	if ok {
		newUsedFor = prevPrices.UsedFor
	}
	// Then add the new recipe to it
	newUsedFor[strconv.Itoa(UserSubmission.RecipeID)] = true
	mainPrices := models.UserPrices{
		ItemID:          UserSubmission.ItemID,
		UsedFor:         newUsedFor,
		MarketItemPrice: UserSubmission.MarketItemPrice,
		MarketAmount:    UserSubmission.MarketAmount,
	}
	newUserStorage.Profits[strconv.Itoa(UserSubmission.RecipeID)] = mainProfits
	newUserStorage.Prices[strconv.Itoa(UserSubmission.ItemID)] = mainPrices

	// Then we add the material prices
	for i := 0; i < len(UserSubmission.IngredientItemID); i++ {
		if UserSubmission.IngredientItemID[i] != 0 {
			matNewUsedFor := make(map[string]bool)
			matItemID := strconv.Itoa(UserSubmission.IngredientItemID[i])
			prevMatPrices, ok := newUserStorage.Prices[matItemID]
			if ok {
				matNewUsedFor = prevMatPrices.UsedFor
			}
			// Then add the new recipe to it
			matNewUsedFor[strconv.Itoa(UserSubmission.RecipeID)] = true

			materialPrices := models.UserPrices{
				ItemID:          UserSubmission.IngredientItemID[i],
				UsedFor:         matNewUsedFor,
				MarketItemPrice: UserSubmission.MarketIngredientPrice[i],
				MarketAmount:    UserSubmission.MarketIngredientAmount[i],
			}
			newUserStorage.Prices[strconv.Itoa(UserSubmission.IngredientItemID[i])] = materialPrices
		}
	}
}
