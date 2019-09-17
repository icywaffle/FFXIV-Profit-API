package models

type Profits struct {
	RecipeID         int     `bson:"RecipeID"`
	ItemID           int     `bson:"ItemID"`
	MaterialCosts    int     `bson:"MaterialCosts"`
	Profits          int     `bson:"Profits"`
	ProfitPercentage float32 `bson:"ProfitPercentage"`
	Added            int64   `bson:"Added"`
	// These are for the profits page.
	// No need to recall Recipes for this.
	Name              string `bson:"Name"`
	IconID            int    `bson:"IconID"`
	CraftTypeTargetID int    `bson:"CraftTypeTargetID"`
	RecipeLevelTable  struct {
		ClassJobLevel int `bson:"ClassJobLevel"`
		Stars         int `bson:"Stars"`
	} `bson:"RecipeLevelTable"`
}

type Recipes struct {
	Name               string `bson:"Name" json:"Name"`
	IconID             int    `bson:"IconID" json:"IconID"`
	ItemResultTargetID int    `bson:"ItemID" json:"ItemResultTargetID"`
	ID                 int    `bson:"RecipeID" json:"ID"`
	CraftTypeTargetID  int    `bson:"CraftTypeTargetID" json:"CraftTypeTargetID"`
	RecipeLevelTable   struct {
		ClassJobLevel int `bson:"ClassJobLevel" json:"ClassJobLevel"`
		Stars         int `bson:"Stars" json:"Stars"`
	} `bson:"RecipeLevelTable" json:"RecipeLevelTable"`
	AmountResult      int        `bson:"AmountResult" json:"AmountResult"`
	IngredientID      [10]int    `bson:"IngredientID"`
	IngredientIconID  [10]int    `bson:"IngredientIconID"`
	IngredientNames   [10]string `bson:"IngredientNames"`
	IngredientAmounts [10]int    `bson:"IngredientAmount"`
	IngredientRecipes [][]int    `bson:"IngredientRecipes"`
	Added             int64      `bson:"Added"`
}

type Prices struct {
	ItemID     int `bson:"ItemID"`
	Sargatanas struct {
		History []struct {
			Added        int  `json:"Added" bson:"Added"` // XIVAPI added time
			IsHQ         bool `json:"IsHQ" bson:"IsHQ"`
			PricePerUnit int  `json:"PricePerUnit" bson:"PricePerUnit"`
			PriceTotal   int  `json:"PriceTotal" bson:"PriceTotal"`
			PurchaseDate int  `json:"PurchaseDate" bson:"PurchaseDate"`
			Quantity     int  `json:"Quantity" bson:"Quantity"`
		} `json:"History" bson:"History"`
		Prices []struct {
			Added        int  `json:"Added" bson:"Added"`
			IsHQ         bool `json:"IsHQ" bson:"IsHQ"`
			PricePerUnit int  `json:"PricePerUnit" bson:"PricePerUnit"`
			PriceTotal   int  `json:"PriceTotal" bson:"PriceTotal"`
			Quantity     int  `json:"Quantity" bson:"Quantity"`
		} `json:"Prices" bson:"Prices"`
	} `json:"Sargatanas" bson:"Sargatanas"`
	OnMarketboard bool  `bson:"OnMarketboard"`
	Added         int64 `bson:"Added"` // Database added time.
}

type SimplePrices struct {
	ItemID            int
	HistoryPrice      int
	LowestMarketPrice int
	OnMarketboard     bool
	Added             int64
}
