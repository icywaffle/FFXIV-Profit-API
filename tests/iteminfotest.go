package tests

import (
	"marketboard-backend/app/controllers/mongoDB/xivapi"

	"github.com/revel/revel/testing"
)

type ItemInfoTest struct {
	testing.TestSuite
}

func (t *ItemInfoTest) Before() {
	println("Set up")
}

///////////////////////// Mocks //////////////////////////

////////////////////// Unit Tests ////////////////////////
func (t *ItemInfoTest) TestIf_Websiteurl_CreatesValidURL() {
	t.AssertEqual("https://xivapi.com/item/14160", xivapi.Websiteurl(14160, "item")[:29])
}

////////////////// Functional Tests //////////////////////
func (t *ItemInfoTest) TestIf_XIVAPI_RecipeEndpointExists() {
	t.AssertNotEqual(nil, xivapi.ApiConnect(33180, "recipe"))
}

func (t *ItemInfoTest) TestIf_XIVAPI_Item_EndpointExists() {
	t.AssertNotEqual(nil, xivapi.ApiConnect(14160, "item"))
}
func (t *ItemInfoTest) After() {
	println("Tear down")
}
