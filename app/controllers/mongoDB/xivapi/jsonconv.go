package xivapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"marketboard-backend/app/models"
	"marketboard-backend/keys"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Converts Recipe Pages of json, to arrays.

/////////////////Recipe Struct Here//////////////////////////

// Unfortunately, this is the state of the API, and what it gives us...
type AmountIngredient struct {
	//The outer values
	AmountIngredient0 int `json:"AmountIngredient0"`
	AmountIngredient1 int `json:"AmountIngredient1"`
	AmountIngredient2 int `json:"AmountIngredient2"`
	AmountIngredient3 int `json:"AmountIngredient3"`
	AmountIngredient4 int `json:"AmountIngredient4"`
	AmountIngredient5 int `json:"AmountIngredient5"`
	AmountIngredient6 int `json:"AmountIngredient6"`
	AmountIngredient7 int `json:"AmountIngredient7"`
	AmountIngredient8 int `json:"AmountIngredient8"`
	AmountIngredient9 int `json:"AmountIngredient9"`
}

type ItemIngredient struct {
	ItemIngredient0 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient0"`
	ItemIngredient1 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient1"`
	ItemIngredient2 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient2"`
	ItemIngredient3 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient3"`
	ItemIngredient4 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient4"`
	ItemIngredient5 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconD"`
	} `json:"ItemIngredient5"`
	ItemIngredient6 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient6"`
	ItemIngredient7 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient7"`
	ItemIngredient8 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient8"`
	ItemIngredient9 struct {
		Name   string `json:"Name"`
		ID     int    `json:"ID"`
		IconID int    `json:"IconID"`
	} `json:"ItemIngredient9"`
}

type IngredientRecipe struct {
	ItemIngredientRecipe0 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe9"`
}

type xivapierror struct {
	Error bool `json:"Error"`
}

var Mutex sync.Mutex

// idtype should be "recipe", "item"
// Will return nil if we get an error response.
func ApiConnect(inputid int, idtype string) []byte {
	// MAX Rate limit is 20 Req/s -> 0.05s/Req.
	// Unfortunately, as we are now, we cannot increase this rate limit.
	// Therefore, if multiple threads start calling this ApiConnect,
	// It would be calling with my key!
	// For a more scalable method, we would need to actually just send a POST of the payload.
	Mutex.Lock()
	time.Sleep(60 * time.Millisecond)
	byteValue := xivapiconnector(Websiteurl(inputid, idtype))
	fmt.Println("Connected to API")
	Mutex.Unlock()

	// Handles invalid json responses.
	var apierror xivapierror
	json.Unmarshal(byteValue, &apierror)
	if apierror.Error {
		return nil
	} else {
		return byteValue
	}

}

func Websiteurl(inputid int, idtype string) string {
	//Example: https://xivapi.com/Item/14160
	basewebsite := []byte("https://xivapi.com/")
	field := []byte(idtype)
	uniqueID := []byte(strconv.Itoa(inputid))
	completefield := append(field[:], '/')
	userinputurl := append(append(basewebsite[:], completefield[:]...), uniqueID[:]...)

	//Add Authkey to the URL
	authkey := []byte(keys.XivAuthKey)
	websiteurl := append(append(userinputurl[:], '?'), authkey[:]...)

	s := string(websiteurl)
	return s
}

// Directly connects to the API here, and returns a byteValue of the body.
func xivapiconnector(websiteurl string) []byte {

	//What this does, is open the file, and read it
	jsonFile, err := http.Get(websiteurl)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	return byteValue
}

func Jsonitemrecipe(byteValue []byte) models.Recipes {

	// Unmarshal the information into the structs
	var recipes models.Recipes
	json.Unmarshal(byteValue, &recipes)

	var amount AmountIngredient
	json.Unmarshal(byteValue, &amount)

	var matitemID ItemIngredient
	json.Unmarshal(byteValue, &matitemID)

	// Create the slices
	amountslice := [10]int{amount.AmountIngredient0,
		amount.AmountIngredient1,
		amount.AmountIngredient2,
		amount.AmountIngredient3,
		amount.AmountIngredient4,
		amount.AmountIngredient5,
		amount.AmountIngredient6,
		amount.AmountIngredient7,
		amount.AmountIngredient8,
		amount.AmountIngredient9}

	matitemIDslice := [10]int{matitemID.ItemIngredient0.ID,
		matitemID.ItemIngredient1.ID,
		matitemID.ItemIngredient2.ID,
		matitemID.ItemIngredient3.ID,
		matitemID.ItemIngredient4.ID,
		matitemID.ItemIngredient5.ID,
		matitemID.ItemIngredient6.ID,
		matitemID.ItemIngredient7.ID,
		matitemID.ItemIngredient8.ID,
		matitemID.ItemIngredient9.ID}

	matitemnameslice := [10]string{matitemID.ItemIngredient0.Name,
		matitemID.ItemIngredient1.Name,
		matitemID.ItemIngredient2.Name,
		matitemID.ItemIngredient3.Name,
		matitemID.ItemIngredient4.Name,
		matitemID.ItemIngredient5.Name,
		matitemID.ItemIngredient6.Name,
		matitemID.ItemIngredient7.Name,
		matitemID.ItemIngredient8.Name,
		matitemID.ItemIngredient9.Name}

	matitemiconslice := [10]int{matitemID.ItemIngredient0.IconID,
		matitemID.ItemIngredient1.IconID,
		matitemID.ItemIngredient2.IconID,
		matitemID.ItemIngredient3.IconID,
		matitemID.ItemIngredient4.IconID,
		matitemID.ItemIngredient5.IconID,
		matitemID.ItemIngredient6.IconID,
		matitemID.ItemIngredient7.IconID,
		matitemID.ItemIngredient8.IconID,
		matitemID.ItemIngredient9.IconID}

	// We need to go through every single possible recipe that can make this item.
	var matrecipeID IngredientRecipe
	json.Unmarshal(byteValue, &matrecipeID)
	matrecipeIDslice := make([][]int, 10)

	//No choice but to unravel for each element, the possible Material Ingredient Recipe IDs 10 times.
	// There is variable length for different elements.
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe0); i++ {
		matrecipeIDslice[0] = append(matrecipeIDslice[0], matrecipeID.ItemIngredientRecipe0[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe1); i++ {
		matrecipeIDslice[1] = append(matrecipeIDslice[1], matrecipeID.ItemIngredientRecipe1[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe2); i++ {
		matrecipeIDslice[2] = append(matrecipeIDslice[2], matrecipeID.ItemIngredientRecipe2[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe3); i++ {
		matrecipeIDslice[3] = append(matrecipeIDslice[3], matrecipeID.ItemIngredientRecipe3[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe4); i++ {
		matrecipeIDslice[4] = append(matrecipeIDslice[4], matrecipeID.ItemIngredientRecipe4[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe5); i++ {
		matrecipeIDslice[5] = append(matrecipeIDslice[5], matrecipeID.ItemIngredientRecipe5[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe6); i++ {
		matrecipeIDslice[6] = append(matrecipeIDslice[6], matrecipeID.ItemIngredientRecipe6[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe7); i++ {
		matrecipeIDslice[7] = append(matrecipeIDslice[7], matrecipeID.ItemIngredientRecipe7[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe8); i++ {
		matrecipeIDslice[8] = append(matrecipeIDslice[8], matrecipeID.ItemIngredientRecipe8[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe9); i++ {
		matrecipeIDslice[9] = append(matrecipeIDslice[9], matrecipeID.ItemIngredientRecipe9[i].ID)
	}

	// These are custom things that we can add to the Recipes documents
	recipes.IngredientID = matitemIDslice
	recipes.IngredientIconID = matitemiconslice
	recipes.IngredientNames = matitemnameslice
	recipes.IngredientAmounts = amountslice
	recipes.IngredientRecipes = matrecipeIDslice
	return recipes
}

func Jsonprices(byteValue []byte) models.Prices {

	var prices models.Prices
	json.Unmarshal(byteValue, &prices)

	return prices

}
