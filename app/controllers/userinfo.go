package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marketboard-backend/app/models"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type UserInfo struct {
	*revel.Controller
}

func oAuth2Discord(AccessToken string) ([]byte, int) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	bearer := fmt.Sprintf("Bearer %s", AccessToken)
	request.Header.Add("Authorization", bearer)

	response, _ := client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)

	return result, response.StatusCode
}

// Given a POST request with just an access token from Discord,
// It stores you into a session, which is really just a cookie.
func (c UserInfo) Login(DiscordToken *models.DiscordToken) revel.Result {
	info, _ := c.Session.Get("DiscordUserID")
	var DiscordUser models.DiscordUser
	if info == nil {
		userbytevalue, StatusCode := oAuth2Discord(DiscordToken.AccessToken)
		json.Unmarshal(userbytevalue, &DiscordUser)

		// If we have an invalid status code, then that means we don't have the right
		// access token. So return.
		if StatusCode != 200 && StatusCode != 201 {
			c.Response.Status = StatusCode
			return c.Render()
		}
		// Assign to the session, the discorduser ID.
		// If we've reached here, that must mean we've properly authenticated.
		c.Session["DiscordUserID"] = DiscordUser.ID
	}
	c.Response.Status = 201
	return c.Render()
}

// Given a POST request with UserSubmission data,
// It handles the data by storing info to the Database
func (c UserInfo) Store(UserSubmission *models.UserSubmission) revel.Result {

	// AUTHENTICATION
	// Checks if there is a session for this user.
	userID, _ := c.Session.Get("DiscordUserID")
	if userID == nil {
		// Forbidden
		c.Response.Status = 403
		return c.Render()
	}

	// STORAGE
	// Adds or updates a user's storage.
	UserItemStorage := UserStorageCollection.FindUserItemStorage(userID.(string))
	if UserItemStorage == nil {
		UserStorageCollection.InsertNewUserItemStorage(UserSubmission, userID.(string))
	} else {
		UserStorageCollection.AddUserItem(UserItemStorage, userID.(string), UserSubmission)
	}

	// 201 - CREATED
	c.Response.Status = 201
	return c.Render()
}

// Given a GET request with a userid and recipeid
// Returns a user's storage document, or nil if there are no user in the database.
// Or returns an empty object if there's no recipe in the database.
func (c UserInfo) Obtain() revel.Result {

	userID, _ := c.Session.Get("DiscordUserID")
	if userID == nil {
		// Forbidden
		c.Response.Status = 403
		return c.Render()
	}

	recipeID := c.Params.Route.Get("recipeid")

	UserItemStorage := UserStorageCollection.FindUserItemStorage(userID.(string))

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
