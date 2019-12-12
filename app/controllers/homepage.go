package controllers

import "github.com/revel/revel"

// Home is the home route of the API
type Home struct {
	*revel.Controller
}

func (c Home) Index() revel.Result {
	return c.RenderTemplate("index.html")
}
