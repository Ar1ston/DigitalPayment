package controllers

import "github.com/revel/revel"

type Error struct {
	*revel.Controller
	Code int
	Msg  string
}

func (c Error) Error(Code int, Msg string) revel.Result {
	return c.Render(Code, Msg)
}
