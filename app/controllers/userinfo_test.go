package controllers_test

import (
	"ffxiv-profit-api/app/controllers"
	"ffxiv-profit-api/app/controllers/mongoDB"
	"ffxiv-profit-api/app/models"
	"ffxiv-profit-api/app/tmp/run"
	"testing"

	"github.com/revel/modules/server-engine/gohttptest/testsuite"
)

// UserInfoTest is a testing Suite for UserInfo
type UserInfoTest struct {
	testsuite.TestSuite
}

// >go test -v -coverprofile=coverage.out ffxiv-profit-api/app/...  -args -revel.importPath=ffxiv-profit-api
/*
	After hours of debugging, suite.Get and any of the HTTP methods provided by just don't work
	You have to use the custom methods, GetCustom, and PostCustom, since the defaults just don't append the right links
*/
func TestMain(m *testing.M) {
	testsuite.RevelTestHelper(m, "dev", run.Run)
}

////////////////// Mocks //////////////////////

// MockUserItemStorage mocks a user's storage in the database
func MockUserItemStorage() *mongoDB.UserItemStorage {
	return &mongoDB.UserItemStorage{
		UserID:  "Test",
		Prices:  make(map[string]models.UserPrices),
		Profits: make(map[string]models.UserProfits),
	}
}

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

////////////////// Unit Tests //////////////////////

// TestIfAddUserInfoToMapChangesInfo checks if we're actually editing our User Info maps
func TestIfAddUserInfoToMapChangesInfo(t *testing.T) {
	suite := testsuite.NewTestSuite(t)
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
	suite.AssertNotEqual(ResultUserItemStorage, *MockUserItemStorage)
}

// TestQuickSortUserProfits actually sorts our response in descending order.
func TestQuickSortUserProfits(t *testing.T) {
	suite := testsuite.NewTestSuite(t)

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
	suite.AssertEqual(sortedMockUserProfits, mockUserProfits)
}

////////////////// Integration Tests //////////////////////
