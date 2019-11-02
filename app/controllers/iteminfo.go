package controllers

import (
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/app/models"

	"github.com/revel/revel"
)

// ItemInfo handles the endpoints for a specific item, without any user information
type ItemInfo struct {
	*revel.Controller
}

// Index shows the single react page for the API.
func (c ItemInfo) Index() revel.Result {

	return c.RenderTemplate("ItemInfo/Index.html")
}

// Obtain grabs all recipe information from XIVAPI, and returns all of the information about the materials and main recipe.
func (c ItemInfo) Obtain(recipeID int) revel.Result {

	var baseinfo mongoDB.Information
	// We have to initialize the maps here, to be able to allow recursive calls.
	var innerinfo mongoDB.InnerInformation
	innerrecipes := make(map[int]*models.Recipes) // Contains the inner recipes for some key = Recipe.ID
	innerinfo.Recipes = innerrecipes
	mongoDB.BaseInformation(DB, recipeID, innerinfo)

	// The baseinfo should also be in the maps themselves.
	baseinfo.Recipes = innerinfo.Recipes[recipeID]

	// We need to render this information as a single JSON object
	jsonObject := make(map[string]interface{})
	jsonObject["MainRecipe"] = baseinfo
	jsonObject["InnerRecipes"] = innerinfo
	return c.RenderJSON(jsonObject)
}
