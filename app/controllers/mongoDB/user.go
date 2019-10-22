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
	Prices  map[string]models.UserPrices
	Profits map[string]models.UserProfits
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
func InsertNewUserItemStorage(userStorage *mongo.Collection, UserSubmission *models.UserSubmission, userID string) {
	newUserStorage := UserItemStorage{
		UserID:  userID,
		Prices:  make(map[string]models.UserPrices),
		Profits: make(map[string]models.UserProfits),
	}
	addUserInfoToMap(&newUserStorage, UserSubmission)

	userStorage.InsertOne(context.TODO(), newUserStorage)
}

// Once we find a specific user's storage, we just add to it and update it.
func AddUserItem(userStorage *mongo.Collection, userItemStorage *UserItemStorage, userID string, UserSubmission *models.UserSubmission) {
	addUserInfoToMap(userItemStorage, UserSubmission)
	filter := bson.M{"userid": userID}
	userStorage.UpdateOne(context.TODO(), filter, bson.D{
		{Key: "$set", Value: userItemStorage},
	})
}

func addUserInfoToMap(newUserStorage *UserItemStorage, UserSubmission *models.UserSubmission) {
	// We add the main recipe first, then it's materials.
	// If the material is craftable, then it'll have it's separate call to this function
	mainProfits := models.UserProfits{
		RecipeID:         UserSubmission.RecipeID,
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
