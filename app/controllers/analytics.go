package controllers

import (
	"context"
	"ffxiv-profit-api/app/models"
	"ffxiv-profit-api/keys"
	"strconv"

	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Analytics for visitors to the API
type Analytics struct {
	*revel.Controller
}

// LogEndpointRequest submits to the database the client IP and endpoint called.
func LogEndpointRequest(request models.EndpointRequest) {
	APIAnalytics.InsertOne(context.TODO(), request)
}

// APIRequestAnalytics returns all IPs and endpoints stored in the database
func APIRequestAnalytics(isAuthorized bool) []*models.EndpointRequest {
	options := options.FindOptions{}
	options.Sort = bson.D{{Key: "requestedtime", Value: 1}}

	limit := int64(1000)
	options.Limit = &limit
	cursor, _ := APIAnalytics.Find(context.Background(), bson.D{}, &options)

	var allRequests []*models.EndpointRequest

	clientObfuscate := make(map[string]string)
	currentNumber := 0
	for cursor.Next(context.TODO()) {
		var singleRequest models.EndpointRequest
		cursor.Decode(&singleRequest)

		// Obfuscate the IP, if not authorized
		if !isAuthorized {
			// User number is actually arbitrary, since it will shift upon new events
			userNumber, ok := clientObfuscate[singleRequest.ClientIP]
			if !ok {
				clientObfuscate[singleRequest.ClientIP] = strconv.Itoa(currentNumber)
				userNumber = strconv.Itoa(currentNumber)
				currentNumber++
			}
			singleRequest.ClientIP = "Hidden" + userNumber
		}

		allRequests = append(allRequests, &singleRequest)
	}
	defer cursor.Close(context.TODO())

	return allRequests
}

// RequestAnalytics is the endpoint that a user should call for analytics stored in the database
func (c Analytics) RequestAnalytics() revel.Result {
	// AUTHENTICATION
	// Checks if there is a session for this user.
	userID, _ := c.Session.Get("DiscordUserID")
	if userID == nil {
		// Forbidden
		c.Response.Status = 403
		return c.Render()
	}

	// Only provide this for valid accounts
	allRequests := []*models.EndpointRequest{}

	// Check if the user is an authorized key or not
	isAuthorized := keys.Authorized()[userID.(string)]
	allRequests = APIRequestAnalytics(isAuthorized)

	jsonObject := make(map[string]interface{})
	jsonObject["Analytics"] = allRequests
	return c.RenderJSON(jsonObject)
}
