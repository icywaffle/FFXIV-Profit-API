package mongoDB

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"marketboard-backend/app/controllers/mongoDB/xivapi"
	"marketboard-backend/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// We want to separate the times, just in case we only update one struct.
// Changing these times will allow us to just update our entries accordingly if the
// structs have been changed.
var UpdatedRecipesStructTime = int64(1563260451) // Last Update : 7/16/19 - 12:01AM
var UpdatedPricesStructTime = int64(1563240349)  // Last Update : 7/15/19 - 6:26PM
var UpdatedProfitsStructTime = int64(1563261129) // Last Update : 7/16/19 - 12:12AM

// We don't want users to create a new mutex every time.
var Mutex sync.Mutex

type Collections struct {
	Prices  *mongo.Collection
	Recipes *mongo.Collection
	Profits *mongo.Collection
}

type Information struct {
	Recipes *models.Recipes
	Prices  *models.SimplePrices
	Profits *models.Profits
}

type InnerInformation struct {
	Recipes map[int]*models.Recipes      // Contains the inner recipes for some key = Recipe.ID
	Prices  map[int]*models.SimplePrices // Contains the inner prices for some key =  Item ID
	Profits map[int]*models.Profits      // Contains the profits for the inner recipes for some key = Recipe.Id
}

type Info struct {
	*Information
	*InnerInformation
}

type CollectionHandler interface {
	FindRecipesDocument(recipeID int) (*models.Recipes, bool)
	FindPricesDocument(itemID int) (*models.Prices, bool)
	FindProfitsDocument(recipeID int) (*models.Profits, bool)
	SimplifyPricesDocument(recipeID int) (*models.SimplePrices, bool)
	InsertRecipesDocument(recipeID int) *models.Recipes
	InsertPricesDocument(itemID int) *models.Prices
	InsertProfitsDocument(profits *models.Profits)
}

type ProfitHandler interface {
	ProfitDescCursor() []*models.Profits
}

// Will return false if there's no recipe in the xivapi.
func (coll Collections) FindRecipesDocument(recipeID int) (*models.Recipes, bool) {
	filter := bson.M{"RecipeID": recipeID}
	var result models.Recipes
	err := coll.Recipes.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, false
	}
	return &result, true
}

func (coll Collections) FindPricesDocument(itemID int) (*models.Prices, bool) {
	filter := bson.M{"ItemID": itemID}
	var result models.Prices
	err := coll.Prices.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, false
	}
	return &result, true
}

// SimplifyPricesDocument is also depreciated
// This is used for quick and easy prices
func (coll Collections) SimplifyPricesDocument(itemID int) (*models.SimplePrices, bool) {
	now := time.Now()
	return &models.SimplePrices{
		ItemID:            itemID,
		HistoryPrice:      0,
		LowestMarketPrice: 0,
		OnMarketboard:     false,
		Added:             now.Unix(),
	}, true
}

// This is used for stronger analysis functions, where we want to see trends etc.
func (coll Collections) FindProfitsDocument(recipeID int) (*models.Profits, bool) {
	filter := bson.M{"RecipeID": recipeID}
	var result models.Profits
	err := coll.Profits.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, false
	}
	return &result, true
}

// Will insert a document, or update it if it's already in the collection.
func (coll Collections) InsertRecipesDocument(recipeID int) *models.Recipes {
	byteValue := xivapi.ApiConnect(recipeID, "recipe")
	// We don't want to be inserting nil values into our xivapi.
	if byteValue != nil {
		result := xivapi.Jsonitemrecipe(byteValue)
		// These variables are not in the json file.
		now := time.Now()
		result.Added = now.Unix()
		// Testing if there's an entry in the DB
		filter := bson.M{"RecipeID": recipeID}

		var options options.CountOptions
		options.SetLimit(1)
		findcount, _ := coll.Recipes.CountDocuments(context.TODO(), filter, &options)
		if findcount < 1 {
			coll.Recipes.InsertOne(context.TODO(), result)
			fmt.Println("Inserted Recipe into Database: ", result.ID)
		} else {
			coll.Recipes.UpdateOne(context.TODO(), filter, bson.D{
				{Key: "$set", Value: result},
			})
			fmt.Println("Updated Item into Recipe Collection :", result.ID)
		}

		return &result
	} else {
		return nil
	}

}

// InsertPrices is Depreciated.
func (coll Collections) InsertPricesDocument(itemID int) *models.Prices {
	now := time.Now()
	sarg := models.Prices{
		ItemID: itemID,
	}
	sarg.Sargatanas.History = []models.History{
		models.History{
			Added:        int(now.Unix()),
			IsHQ:         false,
			PricePerUnit: 0,
			PriceTotal:   0,
			PurchaseDate: 0,
			Quantity:     0,
		},
	}
	sarg.Sargatanas.MarketPrices = []models.MarketPrices{
		models.MarketPrices{
			Added:        int(now.Unix()),
			IsHQ:         false,
			PricePerUnit: 0,
			PriceTotal:   0,
			Quantity:     0,
		},
	}

	return &sarg

}

// Uses the Recipes and Prices from Information, and returns a Profit model.
// Will require profits from the map if the recipe depends on recipes.
func FillProfitsDocument(recipeID int, info InnerInformation) *models.Profits {
	var profits models.Profits

	recipedoc := info.Recipes[recipeID]
	pricesdoc := info.Prices[recipedoc.ItemResultTargetID]

	profits.RecipeID = recipeID
	profits.ItemID = recipedoc.ItemResultTargetID

	// Since we already have the recipe at hand, might as well append it to the profits information
	// This allows to reduce the amount of calls to the database when clicking on profits page.
	profits.Name = recipedoc.Name
	profits.IconID = recipedoc.IconID
	profits.CraftTypeTargetID = recipedoc.CraftTypeTargetID
	profits.RecipeLevelTable.ClassJobLevel = recipedoc.RecipeLevelTable.ClassJobLevel
	profits.RecipeLevelTable.Stars = recipedoc.RecipeLevelTable.Stars

	var materialcost int
	for i := 0; i < len(recipedoc.IngredientID); i++ {
		if recipedoc.IngredientID[i] != 0 {
			// The top of the stack should have no ingredient recipes.
			if recipedoc.IngredientRecipes[i] == nil {
				innerpricedoc := info.Prices[recipedoc.IngredientID[i]]
				materialcost += innerpricedoc.LowestMarketPrice * recipedoc.IngredientAmounts[i]
			} else {
				// If we do have recipes, it should already be defined in the map.
				// For now, we'll do calculations based off of the lowest one
				var innerprofitdoc *models.Profits
				var lowestmaterialcost int
				for j := 0; j < len(recipedoc.IngredientRecipes[i]); j++ {
					tempinnerprofitdoc := info.Profits[recipedoc.IngredientRecipes[i][j]]
					// If lowestmaterialcost is zero, it must mean that we haven't initialized it witha price yet.
					if lowestmaterialcost == 0 || tempinnerprofitdoc.MaterialCosts < lowestmaterialcost {
						innerprofitdoc = info.Profits[recipedoc.IngredientRecipes[i][j]]
					}
				}

				materialcost += innerprofitdoc.MaterialCosts * recipedoc.IngredientAmounts[i]
			}
		}
	}
	profits.MaterialCosts = materialcost

	if pricesdoc.OnMarketboard {
		profits.Profits = pricesdoc.LowestMarketPrice - materialcost
	} else {
		profits.Profits = pricesdoc.HistoryPrice - materialcost
	}

	// Our profit depends on how much money we've spent going into it.
	// Division := 0.1234321
	// Fourinteger := 1234 (Rounded The Final Digit)
	// Profit Percentage := 12.34%
	fourinteger := math.Round(float64(profits.Profits) / float64(materialcost) * 10000)
	// If we have a divide by zero case, we just want to set it to zero then.
	if materialcost == 0 {
		fourinteger = 0
	}
	profits.ProfitPercentage = float32(fourinteger / 100)

	now := time.Now()
	unixtimenow := now.Unix()

	profits.Added = unixtimenow

	return &profits

}
func (coll Collections) InsertProfitsDocument(profits *models.Profits) {
	filter := bson.M{"RecipeID": profits.RecipeID}

	var options options.CountOptions
	options.SetLimit(1)
	findcount, _ := coll.Profits.CountDocuments(context.TODO(), filter, &options)
	if findcount < 1 {
		coll.Profits.InsertOne(context.TODO(), profits)
		fmt.Println("Inserted Profits into Database: ", profits.RecipeID)
	} else {
		coll.Profits.UpdateOne(context.TODO(), filter, bson.D{
			{Key: "$set", Value: profits},
		})
		fmt.Println("Updated Item into Profit Collection :", profits.RecipeID)
	}
}

// Gives a Descending Sorted Array, of 20 items with the most profit from the DB
func (coll Collections) ProfitDescCursor() []*models.Profits {
	options := options.FindOptions{}
	options.Sort = bson.D{{Key: "ProfitPercentage", Value: -1}}
	// We can set this to be bigger later in the future
	limit := int64(20)
	options.Limit = &limit
	cursor, _ := coll.Profits.Find(context.Background(), bson.D{}, &options)

	var allprofits []*models.Profits
	for cursor.Next(context.TODO()) {
		var tempprofits models.Profits
		cursor.Decode(&tempprofits)

		allprofits = append(allprofits, &tempprofits)
	}
	defer cursor.Close(context.TODO())

	return allprofits

}

// Uses recursion to fill the Information maps and inner information.
// A recipe w/ len(IngredientRecipes) = 0, should be at the top of the stack.
// Will handle if there are no items in the xivapi.
// Will also handle struct updates, if you've updated the struct times ontop of xivapi.go.
func BaseInformation(collections CollectionHandler, recipeID int, info InnerInformation) {

	// Adds a base recipe to the map
	baserecipe, indatabase := collections.FindRecipesDocument(recipeID)

	// People can skip the locks if they don't need to insert. (They've already found a document in the database)
	if !indatabase || baserecipe.Added < UpdatedRecipesStructTime {
		Mutex.Lock()
		// Force a recheck for those threads that were waiting on another that was already inserting the same information.
		baserecipe, indatabase = collections.FindRecipesDocument(recipeID)
		if !indatabase || baserecipe.Added < UpdatedRecipesStructTime {
			baserecipe = collections.InsertRecipesDocument(recipeID)
		}
		Mutex.Unlock()
	}
	info.Recipes[recipeID] = baserecipe

	// Finds the prices for the base item of a recipe.
	info.Prices[baserecipe.ItemResultTargetID], indatabase = collections.SimplifyPricesDocument(baserecipe.ItemResultTargetID)

	if !indatabase || info.Prices[baserecipe.ItemResultTargetID].Added < UpdatedPricesStructTime {
		Mutex.Lock()
		info.Prices[baserecipe.ItemResultTargetID], indatabase = collections.SimplifyPricesDocument(baserecipe.ItemResultTargetID)
		if !indatabase || info.Prices[baserecipe.ItemResultTargetID].Added < UpdatedPricesStructTime {
			// It means that the prices are actually not in the database, so we just need to find them.
			collections.InsertPricesDocument(baserecipe.ItemResultTargetID)
			// This also means that we can simplify it.
			info.Prices[baserecipe.ItemResultTargetID], _ = collections.SimplifyPricesDocument(baserecipe.ItemResultTargetID)
		}
		Mutex.Unlock()
	}

	// Also adds all the ingredients prices of current recipe into the map.
	for i := 0; i < len(baserecipe.IngredientID); i++ {
		// Zero are invalid ingredients.
		if baserecipe.IngredientID[i] != 0 {
			info.Prices[baserecipe.IngredientID[i]], indatabase = collections.SimplifyPricesDocument(baserecipe.IngredientID[i])

			if !indatabase || info.Prices[baserecipe.IngredientID[i]].Added < UpdatedPricesStructTime {
				Mutex.Lock()
				info.Prices[baserecipe.IngredientID[i]], indatabase = collections.SimplifyPricesDocument(baserecipe.IngredientID[i])
				if !indatabase || info.Prices[baserecipe.IngredientID[i]].Added < UpdatedPricesStructTime {
					collections.InsertPricesDocument(baserecipe.IngredientID[i])

					info.Prices[baserecipe.IngredientID[i]], _ = collections.SimplifyPricesDocument(baserecipe.IngredientID[i])
				}
				Mutex.Unlock()
			}

		}

	}

	// Recursively call using the inner recipes (if they exist), to add more recipes and prices to our map
	for i := 0; i < len(baserecipe.IngredientRecipes); i++ {
		if baserecipe.IngredientRecipes[i] != nil {
			// Adds all the recipes to the map.
			for j := 0; j < len(baserecipe.IngredientRecipes[i]); j++ {
				BaseInformation(collections, baserecipe.IngredientRecipes[i][j], info)
			}

		}
	}

	// All of our recipes are in the stack. The top of the stack will reach here first and fill our profitmaps.
	baseprofit, indatabase := collections.FindProfitsDocument(recipeID)
	if !indatabase || baseprofit.Added < UpdatedProfitsStructTime {
		Mutex.Lock()
		baseprofit, indatabase = collections.FindProfitsDocument(recipeID)
		if !indatabase || baseprofit.Added < UpdatedProfitsStructTime {
			baseprofit = FillProfitsDocument(recipeID, info)
			collections.InsertProfitsDocument(baseprofit)
		}
		Mutex.Unlock()
	}
	info.Profits[recipeID] = baseprofit
}

func ProfitInformation(profit ProfitHandler) []*models.Profits {

	return profit.ProfitDescCursor()
}
