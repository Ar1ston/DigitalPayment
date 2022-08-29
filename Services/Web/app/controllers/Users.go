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

type requestRemoveUser struct {
	Id uint64 `json:"id"`
}
type respRemoveUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

type requestChangeLevelUser struct {
	Id    int64 `json:"id"`
	Level int8  `json:"level"`
}
type respChangeLevelUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (c Users) Users() revel.Result {
	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}
	var respService respGetUsers

	err := NATS.RequestToNats("Users", "Web", "GetUsers", []byte(""), &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	var usrs []user
	usrs = respService.Users

	return c.Render(usrs)
}
func (c Users) User(id int) revel.Result {
	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}
	var reqService requestGetUser
	reqService.Id = uint64(id)
	var respService respGetUser

	err := NATS.RequestToNats("Users", "Web", "GetUser", &reqService, &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	Login := respService.Login
	Name := respService.Name
	Level := respService.Level

	return c.Render(id, Name, Login, Level)
}
func (c Users) Remove(id int) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	var reqService requestRemoveUser
	reqService.Id = uint64(id)

	var respService respRemoveUser

	err := NATS.RequestToNats("Users", "Web", "RemoveUser", &reqService, &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	return c.Redirect(Users.Users)
}
func (c Users) ChangeLevel(id int, level int) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	var reqService requestChangeLevelUser
	reqService.Id = int64(id)
	reqService.Level = int8(level)
	var respService respChangeLevelUser

	err := NATS.RequestToNats("Users", "Web", "ChangeLevelUser", &reqService, &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	return c.Redirect(Users.Users)
}
