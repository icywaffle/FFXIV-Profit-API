package models

// UserSubmission is an object that holds the payload that is necessary to store in the database.
type UserSubmission struct {
	UserID           string
	RecipeID         int
	ItemID           int
	Profits          int
	ProfitPercentage int
	MarketItemPrice  int
	MarketAmount     int

	IngredientItemID       []int
	MarketIngredientPrice  []int
	MarketIngredientAmount []int
}

// UserPrices are the POSTed user prices for a specific item
type UserPrices struct {
	ItemID          int
	UsedFor         map[string]bool
	MarketItemPrice int
	MarketAmount    int
}

// UserProfits is the calculated profit for a specific recipe.
type UserProfits struct {
	RecipeID         int
	ItemID           int
	IngredientItemID []int
	Profits          int
	ProfitPercentage int
}
