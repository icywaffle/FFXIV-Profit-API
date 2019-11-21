package controllers

import (
	"context"
	"ffxiv-profit-api/app/models"
)

// LogEndpointRequest submits to the database the client IP and endpoint called.
func LogEndpointRequest(request models.EndpointRequest) {
	APIAnalytics.InsertOne(context.TODO(), request)
}
