package controllers

import (
	"Web/conf/NATS"
	"Web/conf/crypt"
	"fmt"
	"github.com/revel/revel"
)

type Login struct {
	*revel.Controller
}
type requestCheckUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type respCheckUser struct {
	Id    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Level uint64 `json:"level"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (c Login) Login(login string, password string) revel.Result {
	if c.Request.Method == "POST" {
		var reqService requestCheckUser
		reqService.Login = login

		cryptPassword, err := crypt.Aes_encrypt(password)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}
		reqService.Password = cryptPassword
		var respService respCheckUser

		err = NATS.RequestToNats("Users", "Web", "CheckUser", &reqService, &respService)
		if err != nil {
			return c.Render()
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Render()
		}

		c.Session["login"] = respService.Login
		c.Session["name"] = respService.Name
		c.Session["level"] = respService.Level
		c.Session["id"] = respService.Id

		return c.Redirect(Books.Books)
	}
	delete(c.Session, "login")
	delete(c.Session, "level")
	delete(c.Session, "name")
	delete(c.Session, "id")
	return c.Render()
}
