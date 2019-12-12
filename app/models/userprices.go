package models

// UserSubmission is an object that holds the payload that is necessary to store in the database.
type UserSubmission struct {
	// Profit Information
	UserID           string
	RecipeID         int
	ItemID           int
	ItemName         string
	IconID           int
	MaterialCosts    int
	Profits          int
	ProfitPercentage int
	MarketItemPrice  int
	MarketAmount     int

	// Price Information
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
	Added            int64
	RecipeID         int
	ItemID           int
	ItemName         string
	IconID           int
	IngredientItemID []int
	MaterialCosts    int
	Profits          int
	ProfitPercentage int
}
