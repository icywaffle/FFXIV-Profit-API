package models

type UserPrice struct {
	UserID                string
	RecipeID              int
	MarketItemPrice       int
	MarketIngredientPrice []int
	MarketAmount          []int
}
