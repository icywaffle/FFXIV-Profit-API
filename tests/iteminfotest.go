package tests

import (
	"marketboard-backend/app/controllers/mongoDB/xivapi"

	"github.com/revel/revel/testing"
)

// ItemInfoTest is the testing suite for ItemInfo
type ItemInfoTest struct {
	testing.TestSuite
}

// Before tests initializes
func (t *ItemInfoTest) Before() {
	println("Set up")
}

///////////////////////// Mocks //////////////////////////

////////////////////// Unit Tests ////////////////////////

// TestIfWebsiteurlCreatesValidURL checks if we're actually creating the correct URL.
func (t *ItemInfoTest) TestIfWebsiteurlCreatesValidURL() {
	t.AssertEqual("https://xivapi.com/item/14160", xivapi.Websiteurl(14160, "item")[:29])
}

////////////////// Functional Tests //////////////////////

// TestIfXIVAPIRecipeEndpointExists checks if our Recipe endpoint is still valid.
func (t *ItemInfoTest) TestIfXIVAPIRecipeEndpointExists() {
	t.AssertNotEqual(nil, xivapi.APIConnect(33180, "recipe"))
}

// TestIfXIVAPIItemEndpointExists checks if our Item endpoint is still valid.
func (t *ItemInfoTest) TestIfXIVAPIItemEndpointExists() {
	t.AssertNotEqual(nil, xivapi.APIConnect(14160, "item"))
}

// After tests finishes
func (t *ItemInfoTest) After() {
	println("Tear down")
}
