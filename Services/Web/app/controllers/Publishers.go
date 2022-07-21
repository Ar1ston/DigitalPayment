package controllers

import "github.com/revel/revel"

type Publishers struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Publishers) Publishers() revel.Result {
	return c.Render()
}
