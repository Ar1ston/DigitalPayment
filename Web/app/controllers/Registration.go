package controllers

import (
	"Web/conf/NATS"
	"Web/conf/crypt"
	"fmt"
	"github.com/revel/revel"
)

type Registration struct {
	*revel.Controller
	FirstName   string
	LastName    string
	Description string
}
type requestCreateUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type responseCreateUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (c Registration) Registration(login string, password string, name string) revel.Result {
	if c.Request.Method == "POST" {
		var reqService requestCreateUser
		reqService.Login = login

		cryptPassword, err := crypt.Aes_encrypt(password)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}

		reqService.Password = cryptPassword
		reqService.Name = name
		var respService responseCreateUser

		err = NATS.RequestToNats("Users", "Web", "CreateUser", &reqService, &respService)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
		}

		return c.Redirect(Login.Login)
	}
	delete(c.Session, "login")
	delete(c.Session, "level")
	delete(c.Session, "name")
	delete(c.Session, "id")
	return c.Render()
}
