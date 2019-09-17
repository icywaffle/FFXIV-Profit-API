package controllers

import (
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/app/models"

	"github.com/revel/revel"
)

type ItemInfo struct {
	*revel.Controller
}

func (c ItemInfo) Index() revel.Result {

	return c.RenderTemplate("Result/Index.html")
}

func (c ItemInfo) Obtain(recipeID int) revel.Result {

	var baseinfo mongoDB.Information
	// We have to initialize the maps here, to be able to allow recursive calls.
	var innerinfo mongoDB.InnerInformation
	innerrecipes := make(map[int]*models.Recipes)           // Contains the inner recipes for some key = Recipe.ID
	innersimpleprices := make(map[int]*models.SimplePrices) // Contains the inner prices for some key =  Item ID
	innerprofits := make(map[int]*models.Profits)
	innerinfo.Recipes = innerrecipes
	innerinfo.Prices = innersimpleprices
	innerinfo.Profits = innerprofits
	mongoDB.BaseInformation(DB, recipeID, innerinfo)

	// The baseinfo should also be in the maps themselves.
	baseinfo.Recipes = innerinfo.Recipes[recipeID]
	baseinfo.Prices = innerinfo.Prices[baseinfo.Recipes.ItemResultTargetID]
	baseinfo.Profits = innerinfo.Profits[recipeID]

	// We need to render this information as a single JSON object
	jsonObject := make(map[string]interface{})
	jsonObject["MainRecipe"] = baseinfo
	jsonObject["InnerRecipes"] = innerinfo

	return c.RenderJSON(jsonObject)
}

func (c ItemInfo) Profit() revel.Result {
	profitpercentage := mongoDB.ProfitInformation(DB)

	jsonObject := make(map[string]interface{})
	jsonObject["Profits"] = profitpercentage

	return c.RenderJSON(jsonObject)
}
