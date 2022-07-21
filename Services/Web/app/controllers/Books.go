package controllers

import "github.com/revel/revel"

type Books struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Books) Books() revel.Result {
	return c.Render()
}
