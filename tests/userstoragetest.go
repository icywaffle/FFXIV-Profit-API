package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marketboard-backend/app/models"

	"github.com/revel/revel/testing"
)

type UserInfoTest struct {
	testing.TestSuite
}

func (t *UserInfoTest) Before() {
	println("Set up")
}

func (t *UserInfoTest) TestIfPostRequestSucceeded() {
	data := models.UserPrice{
		UserID:                0001,
		RecipeID:              33180,
		MarketItemPrice:       100,
		MarketIngredientPrice: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		MarketAmount:          []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	dataByte, _ := json.Marshal(data)
	dataReader := bytes.NewReader(dataByte)
	t.Post("/userinfo/0001", "application/json", dataReader)
	bodyBytes, _ := ioutil.ReadAll(t.Response.Body)
	fmt.Println(string(bodyBytes))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *UserInfoTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *UserInfoTest) After() {
	println("Tear down")
}
