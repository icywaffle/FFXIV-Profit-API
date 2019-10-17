package controllers

import (
	"fmt"
	"marketboard-backend/app/models"

	"github.com/revel/revel"
)

type UserInfo struct {
	*revel.Controller
}

func (c UserInfo) Store(userPrice *models.UserPrice) revel.Result {
	fmt.Println(userPrice)

	jsonObject := make(map[string]interface{})
	jsonObject["message"] = "success"
	return c.RenderJSON(jsonObject)
}
