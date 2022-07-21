package controllers

import "github.com/revel/revel"

type Registration struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Registration) Registration() revel.Result {
	return c.Render()
}
