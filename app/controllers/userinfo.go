package controllers

import (
	"marketboard-backend/app/controllers/mongoDB"
	"marketboard-backend/app/models"

	"github.com/revel/revel"
)

type UserInfo struct {
	*revel.Controller
}

func (c UserInfo) Store(userPrice *models.UserPrice) revel.Result {
	UserItemStorage := mongoDB.FindUserItemStorage(UserStorage, userPrice.UserID)
	if UserItemStorage == nil {
		mongoDB.InsertNewUserItemStorage(UserStorage, userPrice.UserID)
	} else {
		mongoDB.AddUserItem(UserStorage, UserItemStorage, userPrice.UserID, userPrice)
	}
	jsonObject := make(map[string]interface{})
	jsonObject["message"] = "success"
	return c.RenderJSON(jsonObject)
}
