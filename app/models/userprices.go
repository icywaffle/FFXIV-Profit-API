package models

type UserPrice struct {
	UserID                int
	RecipeID              int
	MarketItemPrice       int
	MarketIngredientPrice []int
	MarketAmount          []int
}
