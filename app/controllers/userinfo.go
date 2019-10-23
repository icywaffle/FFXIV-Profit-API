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

	// We need to find all the ingredients for a specific recipe
	// This is so that we return all the prices that are relevant to a specific recipe
	RecipeID, _ := strconv.Atoi(recipeID)
	RecipeDocument, inDB := DB.FindRecipesDocument(RecipeID)
	ingredients := [10]int{}
	if inDB {
		ingredients = RecipeDocument.IngredientID
	}

	jsonObject := make(map[string]interface{})
	// If we don't have an object, just autofill to zero for now.
	if UserItemStorage != nil {
		RecipeProfits, ok := UserItemStorage.Profits[recipeID]
		if ok {
			jsonObject["UserProfits"] = RecipeProfits
		}
		// We just want specific prices for items related to our recipe.
		recipeItems := make(map[int]models.UserPrices)
		// Don't forget we also want price info for the item being made by the recipe too
		itemID := RecipeDocument.ItemResultTargetID
		RecipePrices, ok := UserItemStorage.Prices[strconv.Itoa(itemID)]
		if ok {
			recipeItems[itemID] = RecipePrices
		}

		for i := 0; i < len(ingredients); i++ {
			if ingredients[i] != 0 {
				MaterialPrices, ok := UserItemStorage.Prices[strconv.Itoa(ingredients[i])]
				if ok {
					recipeItems[ingredients[i]] = MaterialPrices
				}

			}
		}

		if len(recipeItems) > 0 {
			jsonObject["UserPrices"] = recipeItems
		}

	}

	return c.RenderJSON(jsonObject)
}
