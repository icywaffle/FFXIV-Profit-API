package mongoDB

import (
	"context"
	"fmt"
	"sync"
	"time"

	"marketboard-backend/app/controllers/mongoDB/xivapi"
	"marketboard-backend/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// We don't want users to create a new mutex every time.
var Mutex sync.Mutex

type Collections struct {
	Recipes *mongo.Collection
}

type Information struct {
	Recipes *models.Recipes
}

type InnerInformation struct {
	Recipes map[int]*models.Recipes // Contains the inner recipes for some key = Recipe.ID
}

type Info struct {
	*Information
	*InnerInformation
}

type CollectionHandler interface {
	FindRecipesDocument(recipeID int) (*models.Recipes, bool)
	InsertRecipesDocument(recipeID int) *models.Recipes
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

// Uses recursion to fill the Information maps and inner information.
// A recipe w/ len(IngredientRecipes) = 0, should be at the top of the stack.
// Will handle if there are no items in the xivapi.
// Will also handle struct updates, if you've updated the struct times ontop of xivapi.go.
func BaseInformation(collections CollectionHandler, recipeID int, info InnerInformation) {

	// Adds a base recipe to the map
	baserecipe, indatabase := collections.FindRecipesDocument(recipeID)

	// People can skip the locks if they don't need to insert. (They've already found a document in the database)
	if !indatabase {
		Mutex.Lock()
		// Force a recheck for those threads that were waiting on another that was already inserting the same information.
		baserecipe, indatabase = collections.FindRecipesDocument(recipeID)
		if !indatabase {
			baserecipe = collections.InsertRecipesDocument(recipeID)
		}
		Mutex.Unlock()
	}
	info.Recipes[recipeID] = baserecipe

	// Recursively call using the inner recipes (if they exist), to add more recipes and prices to our map
	for i := 0; i < len(baserecipe.IngredientRecipes); i++ {
		if baserecipe.IngredientRecipes[i] != nil {
			// Adds all the recipes to the map.
			for j := 0; j < len(baserecipe.IngredientRecipes[i]); j++ {
				BaseInformation(collections, baserecipe.IngredientRecipes[i][j], info)
			}

		}
	}

}
