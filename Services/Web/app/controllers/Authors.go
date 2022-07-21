package controllers

import "github.com/revel/revel"

type Authors struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Authors) Authors() revel.Result {
	return c.Render()
}
