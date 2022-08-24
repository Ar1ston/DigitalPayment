package controllers

import (
	"Web/conf/NATS"
	"fmt"
	"github.com/revel/revel"
)

type Authors struct {
	*revel.Controller
}

type author struct {
	Id        uint64
	FirstName string
	LastName  string
	//Description string
}
type respAuthors struct {
	Authors []author
	Errno   uint64 `json:"errno"`
	Error   string `json:"error,omitempty"`
}
type respAuthor struct {
	Id          uint64 `json:"Id"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}
type respRemoveAuthor struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}
type requestGetAuthor struct {
	Id uint64 `json:"Id"`
}
type requestRemoveAuthor struct {
	Id uint64 `json:"Id"`
}
type requestChangeAuthor struct {
	Id          uint64 `json:"Id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type requestCreateAuthor struct {
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c Authors) Authors() revel.Result {
	var respService respAuthors
	err := NATS.RequestToNats("Authors", "Web", "GetAuthors", []byte(""), &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	var auth []author
	auth = respService.Authors

	return c.Render(auth)
}
func (c Authors) Author(id int) revel.Result {
	var reqService requestGetAuthor
	reqService.Id = uint64(id)
	var respService respAuthor

	err := NATS.RequestToNats("Authors", "Web", "GetAuthor", &reqService, &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	firstName := respService.FirstName
	lastName := respService.LastName
	desc := respService.Description

	return c.Render(id, firstName, lastName, desc)
}
func (c Authors) Remove(id int) revel.Result {
	var reqService requestRemoveAuthor
	reqService.Id = uint64(id)

	var respService respRemoveAuthor

	err := NATS.RequestToNats("Authors", "Web", "RemoveAuthor", &reqService, &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	return c.Redirect(Authors.Authors)
}
func (c Authors) Change(id int, FirstName string, LastName string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		var reqService requestChangeAuthor
		reqService.Id = uint64(id)
		reqService.First_name = FirstName
		reqService.Last_name = LastName
		reqService.Description = Description
		var respService respRemoveAuthor

		err := NATS.RequestToNats("Authors", "Web", "ChangeAuthor", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		return c.Redirect(Authors.Authors)
	} else {

		var reqService requestGetAuthor
		reqService.Id = uint64(id)
		var respService respAuthor

		err := NATS.RequestToNats("Authors", "Web", "GetAuthor", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		firstName := respService.FirstName
		lastName := respService.LastName
		desc := respService.Description

		return c.Render(id, firstName, lastName, desc)
	}
}
func (c Authors) Create(FirstName string, LastName string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		var reqService requestCreateAuthor
		reqService.First_name = FirstName
		reqService.Last_name = LastName
		reqService.Description = Description

		var respService respRemoveAuthor

		err := NATS.RequestToNats("Authors", "Web", "CreateAuthor", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		return c.Redirect(Authors.Authors)
	}
	return c.Render()
}
