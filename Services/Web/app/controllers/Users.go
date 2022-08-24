package controllers

import (
	"Web/conf/NATS"
	"fmt"
	"github.com/revel/revel"
)

type Users struct {
	*revel.Controller
}
type user struct {
	Id    uint64
	Name  string
	Login string
	Level uint64
}
type respGetUsers struct {
	Users []user
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

type requestGetUser struct {
	Id uint64 `json:"id"`
}
type respGetUser struct {
	Id    uint64 `json:"Id"`
	Login string `json:"Login"`
	Name  string `json:"Name"`
	Level uint64 `json:"level"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (c Users) Users() revel.Result {
	var respService respGetUsers

	err := NATS.RequestToNats("Users", "Web", "GetUsers", []byte(""), &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	var usrs []user
	usrs = respService.Users

	return c.Render(usrs)
}
func (c Users) User(id int) revel.Result {
	var reqService requestGetUser
	reqService.Id = uint64(id)
	var respService respGetUser

	err := NATS.RequestToNats("Users", "Web", "GetUser", &reqService, &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	Login := respService.Login
	Name := respService.Name
	Level := respService.Level

	return c.Render(Name, Login, Level)
}
