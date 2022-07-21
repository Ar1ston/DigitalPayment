package controllers

import "github.com/revel/revel"

type Login struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}

func (c Login) Login() revel.Result {
	return c.Render()
}
