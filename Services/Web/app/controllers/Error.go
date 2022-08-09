package controllers

import "github.com/revel/revel"

type Error struct {
	*revel.Controller
	Code int
	Msg  string
}

func (c Error) Error() revel.Result {
	return c.Render(c.Code, c.Msg)
}
