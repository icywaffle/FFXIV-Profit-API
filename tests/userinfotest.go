package tests

import (
	"bytes"
	"encoding/json"
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/app/models"

	"github.com/revel/revel/testing"
)

type UserInfoTest struct {
	testing.TestSuite
}

func (t *UserInfoTest) Before() {
	println("Set up")
}

///////////////////////// Mocks //////////////////////////
func MockUserSubmissionData() *models.UserSubmission {
	return &models.UserSubmission{
		UserID:           "Test",
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

func MockUserItemStorage() *mongoDB.UserItemStorage {
	return &mongoDB.UserItemStorage{
		UserID:  "Test",
		Prices:  make(map[string]models.UserPrices),
		Profits: make(map[string]models.UserProfits),
	}
}

////////////////////// Unit Tests ////////////////////////
func (t *UserInfoTest) TestIf_AddUserInfoToMap_ChangesInfo() {
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

////////////////// Functional Tests //////////////////////

// Checks if we can send a proper POST request to the backend server
func (t *UserInfoTest) TestIfPOSTuserinfoSucceeded() {
	dataByte, _ := json.Marshal(MockUserSubmissionData())
	dataReader := bytes.NewReader(dataByte)
	t.Post("/userinfo/Test", "application/json", dataReader)
	t.AssertEqual(t.Response.Status, "200 OK")
}

// Checks if we can send a GET request to the backend server for a user info
func (t *UserInfoTest) TestifGETuserinfoSucceeded() {
	t.Get("/userinfo/Test/recipe/33180")
	t.AssertEqual(t.Response.Status, "200 OK")
}

// We need to unmarshal the response from /userinfo/Test/recipe/:recipeID
type UserInfoResponse struct {
	UserPrices  map[string]models.UserPrices  `json:"UserPrices"`
	UserProfits map[string]models.UserProfits `json:"UserProfits"`
}

// Checks if we can obtain data from the database after posting data
func (t *UserInfoTest) TestIfDatabaseDocumentExistsAfterPOST() {
	// We need to refresh the data just in case
	dataByte, _ := json.Marshal(MockUserSubmissionData())
	dataReader := bytes.NewReader(dataByte)
	t.Post("/userinfo/Test", "application/json", dataReader)

	// Then we can check if the data is not empty
	t.Get("/userinfo/Test/recipe/33180")
	var UserInfoResponse UserInfoResponse
	emptyUserInfoResponse := UserInfoResponse
	json.Unmarshal(t.ResponseBody, &UserInfoResponse)
	t.AssertNotEqual(UserInfoResponse, emptyUserInfoResponse)
}
func (t *UserInfoTest) After() {
	println("Tear down")
}
