package tests

import (
	"bytes"
	"encoding/json"
	"ffxiv-profit-api/app/controllers"
	"ffxiv-profit-api/app/controllers/mongoDB"
	"ffxiv-profit-api/app/models"
	"ffxiv-profit-api/keys"
	"fmt"

	"github.com/revel/revel/testing"
)

// UserInfoTest is a testing Suite for UserInfo
type UserInfoTest struct {
	testing.TestSuite
}

// Before initializes before our tests runs.
func (t *UserInfoTest) Before() {
	fmt.Println("Setup")

	// We need to log in first
	AuthToken := models.DiscordToken{
		AccessToken: keys.TestAuthKey,
	}
	dataByte, _ := json.Marshal(AuthToken)
	dataReader := bytes.NewReader(dataByte)
	t.Post("/api/testlogin", "application/json", dataReader)
	fmt.Println(t.Response.Status)
}

///////////////////////// Mocks //////////////////////////

// MockUserSubmissionData mocks a Payload to the POST /userinfo/ endpoint
func MockUserSubmissionData() *models.UserSubmission {
	return &models.UserSubmission{
		RecipeID:         33180,
		ItemID:           24678,
		Profits:          0,
		ProfitPercentage: 0,
		MarketItemPrice:  100,
		MarketAmount:     1,

		IngredientItemID:       []int{14146, 14155, 14149, 15653, 16733, 0, 0, 0, 14, 17},
		MarketIngredientPrice:  []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		MarketIngredientAmount: []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// MockUserItemStorage mocks a user's storage in the database
func MockUserItemStorage() *mongoDB.UserItemStorage {
	return &mongoDB.UserItemStorage{
		UserID:  "Test",
		Prices:  make(map[string]models.UserPrices),
		Profits: make(map[string]models.UserProfits),
	}
}

////////////////////// Unit Tests ////////////////////////

// TestIfAddUserInfoToMapChangesInfo checks if we are editing our user's storage.
func (t *UserInfoTest) TestIfAddUserInfoToMapChangesInfo() {
	MockUserItemStorage := MockUserItemStorage()
	MockUserSubmissionData := MockUserSubmissionData()
	mongoDB.AddUserInfoToMap(MockUserItemStorage, MockUserSubmissionData)

	// We need to generate what the result would look like
	ResultUserItemStorage := mongoDB.UserItemStorage{
		UserID: "Test",
		// These are the important things that need info
		Prices:  make(map[string]models.UserPrices),
		Profits: make(map[string]models.UserProfits),
	}

	// If it changes, then that should mean that we're editing info
	t.AssertNotEqual(ResultUserItemStorage, *MockUserItemStorage)
}

// TestQuickSortUserProfits actually sorts our response in descending order.
func (t *UserInfoTest) TestQuickSortUserProfits() {
	mockUserProfits := make([]models.UserProfits, 4)

	mockUserProfits[0] = models.UserProfits{
		RecipeID:         1,
		ProfitPercentage: 1,
	}
	mockUserProfits[1] = models.UserProfits{
		RecipeID:         2,
		ProfitPercentage: -100,
	}

	mockUserProfits[2] = models.UserProfits{
		RecipeID:         3,
		ProfitPercentage: 200,
	}
	mockUserProfits[3] = models.UserProfits{
		RecipeID:         4,
		ProfitPercentage: 20,
	}

	sortedMockUserProfits := make([]models.UserProfits, 4)
	// Eyeball sorted
	sortedMockUserProfits[0] = mockUserProfits[2]
	sortedMockUserProfits[1] = mockUserProfits[3]
	sortedMockUserProfits[2] = mockUserProfits[0]
	sortedMockUserProfits[3] = mockUserProfits[1]

	controllers.QuickSortUserProfits(mockUserProfits, 0, len(mockUserProfits)-1)
	t.AssertEqual(sortedMockUserProfits, mockUserProfits)

}

////////////////// Functional Tests //////////////////////

// TestIfPOSTuserinfoSucceeded checks if we can send a proper POST request to the backend server
func (t *UserInfoTest) TestIfPOSTuserinfoSucceeded() {
	dataByte, _ := json.Marshal(MockUserSubmissionData())
	dataReader := bytes.NewReader(dataByte)
	t.Post("/api/userinfo/", "application/json", dataReader)
	t.AssertEqual(t.Response.Status, "201 Created")
}

// TestifGETuserinfoSucceeded checks if we can send a GET request to the backend server for a user info
func (t *UserInfoTest) TestifGETuserinfoSucceeded() {
	t.Get("/api/userinfo/recipe/33180")
	t.AssertEqual(t.Response.Status, "200 OK")
}

// UserInfoResponse is a placeholder to unmarshal the response from /userinfo/Test/recipe/:recipeID
type UserInfoResponse struct {
	UserPrices  map[string]models.UserPrices  `json:"UserPrices"`
	UserProfits map[string]models.UserProfits `json:"UserProfits"`
}

// TestIfDatabaseDocumentExistsAfterPOST checks if we can obtain data from the database after posting data
func (t *UserInfoTest) TestIfDatabaseDocumentExistsAfterPOST() {
	// We need to refresh the data just in case
	dataByte, _ := json.Marshal(MockUserSubmissionData())
	dataReader := bytes.NewReader(dataByte)
	t.Post("/api/userinfo/", "application/json", dataReader)

	// Then we can check if the data is not empty
	t.Get("/api/userinfo/recipe/33180")
	var UserInfoResponse UserInfoResponse
	emptyUserInfoResponse := UserInfoResponse
	json.Unmarshal(t.ResponseBody, &UserInfoResponse)
	t.AssertNotEqual(UserInfoResponse, emptyUserInfoResponse)
}

// After shows when tests are completed
func (t *UserInfoTest) After() {
	fmt.Println("Tear Down")
	// Logout when done
	t.Get("/api/userinfo/logout")
}
