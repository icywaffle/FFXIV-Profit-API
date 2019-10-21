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

func (c UserInfo) Store(userPrice *models.UserPrice) revel.Result {
	UserItemStorage := mongoDB.FindUserItemStorage(UserStorage, userPrice.UserID)
	if UserItemStorage == nil {
		mongoDB.InsertNewUserItemStorage(UserStorage, userPrice, userPrice.UserID)
	} else {
		mongoDB.AddUserItem(UserStorage, UserItemStorage, userPrice.UserID, userPrice)
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
	// If we don't have an object, just autofill to zero for now.
	if UserItemStorage != nil && UserItemStorage.Recipe[recipeID] != nil {
		jsonObject["UserPrice"] = UserItemStorage.Recipe[recipeID]
	} else {
		recipe, _ := strconv.Atoi(recipeID)
		jsonObject["UserPrice"] = models.UserPrice{
			UserID:                userID,
			RecipeID:              recipe,
			MarketItemPrice:       0,
			MarketIngredientPrice: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			MarketAmount:          []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
	}

	return c.RenderJSON(jsonObject)
}
