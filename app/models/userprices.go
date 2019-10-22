package models

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

type UserPrices struct {
	ItemID          int
	UsedFor         map[string]bool
	MarketItemPrice int
	MarketAmount    int
}

type UserProfits struct {
	RecipeID         int
	ItemID           int
	IngredientItemID []int
	Profits          int
	ProfitPercentage int
}
