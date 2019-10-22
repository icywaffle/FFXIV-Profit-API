package controllers

import (
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/app/models"
	"strconv"

	"github.com/revel/revel"
)

type UserInfo struct {
	*revel.Controller
}

func (c UserInfo) Store(UserSubmission *models.UserSubmission) revel.Result {
	UserItemStorage := mongoDB.FindUserItemStorage(UserStorage, UserSubmission.UserID)
	if UserItemStorage == nil {
		mongoDB.InsertNewUserItemStorage(UserStorage, UserSubmission, UserSubmission.UserID)
	} else {
		mongoDB.AddUserItem(UserStorage, UserItemStorage, UserSubmission.UserID, UserSubmission)
	}
	jsonObject := make(map[string]interface{})
	jsonObject["message"] = "success"
	return c.RenderJSON(jsonObject)
}

func (c UserInfo) Obtain() revel.Result {
	userID := c.Params.Route.Get("userid")
	recipeID := c.Params.Route.Get("recipeid")

	UserItemStorage := mongoDB.FindUserItemStorage(UserStorage, userID)

	jsonObject := make(map[string]interface{})

	// IDs of 0 are invalid
	if recipeID == "0" {
		return c.RenderJSON(jsonObject)
	}

	// If we don't have an object, just autofill to zero for now.
	if UserItemStorage != nil {
		jsonObject["UserProfits"] = UserItemStorage.Profits[recipeID]
		// We just want specific prices for items related to our recipe.
		recipeItems := make(map[int]models.UserPrices)
		// Don't forget we also want price info for the item being made by the recipe too
		itemID := UserItemStorage.Profits[recipeID].ItemID
		recipeItems[itemID] = UserItemStorage.Prices[strconv.Itoa(itemID)]

		ingredientItemID := UserItemStorage.Profits[recipeID].IngredientItemID
		for i := 0; i < len(ingredientItemID); i++ {
			if ingredientItemID[i] != 0 {
				recipeItems[ingredientItemID[i]] = UserItemStorage.Prices[strconv.Itoa(ingredientItemID[i])]
			}
		}
		jsonObject["UserPrices"] = recipeItems
	}

	return c.RenderJSON(jsonObject)
}
