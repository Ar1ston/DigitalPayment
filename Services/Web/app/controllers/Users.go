package controllers

import "github.com/revel/revel"

type Users struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Users) Users() revel.Result {
	return c.Render()
}
