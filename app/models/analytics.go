package models

import "time"

type EndpointRequest struct {
	ClientIP      string
	Endpoint      string
	RequestedTime time.Time
}
