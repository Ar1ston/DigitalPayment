package controllers

import (
	"Web/conf/NATS"
	"fmt"
	"github.com/revel/revel"
)

type Publishers struct {
	*revel.Controller
}

type publisher struct {
	Id   uint64
	Name string
}
type respPublishers struct {
	Publishers []publisher
	Errno      uint64 `json:"errno"`
	Error      string `json:"error,omitempty"`
}
type respPublisher struct {
	Id          uint64 `json:"Id"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}
type respRemovePublisher struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}
type requestGetPublisher struct {
	Id uint64 `json:"Id"`
}
type requestRemovePublisher struct {
	Id uint64 `json:"Id"`
}
type requestChangePublisher struct {
	Id          uint64 `json:"Id"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"description,omitempty"`
}
type requestCreatePublisher struct {
	Name        string `json:"Name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c Publishers) Publishers() revel.Result {

	var respService respPublishers

	err := NATS.RequestToNats("Publishers", "Web", "GetPublishers", []byte(""), &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	var pubs []publisher
	pubs = respService.Publishers
	return c.Render(pubs)
}
func (c Publishers) Publisher(id int) revel.Result {

	var reqService requestGetPublisher
	reqService.Id = uint64(id)
	var respService respPublisher

	err := NATS.RequestToNats("Publishers", "Web", "GetPublisher", &reqService, &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	name := respService.Name
	desc := respService.Description
	return c.Render(id, name, desc)
}
func (c Publishers) Remove(id int) revel.Result {

	var reqService requestRemovePublisher
	reqService.Id = uint64(id)

	var respService respPublisher

	err := NATS.RequestToNats("Publishers", "Web", "RemovePublisher", &reqService, &respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	return c.Redirect(Publishers.Publishers)
}
func (c Publishers) Change(id int, Name string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		var reqService requestChangePublisher
		reqService.Id = uint64(id)
		reqService.Name = Name
		reqService.Description = Description
		var respService respRemovePublisher

		err := NATS.RequestToNats("Publishers", "Web", "ChangePublisher", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		return c.Redirect(Publishers.Publishers)
	} else {

		var reqService requestGetPublisher
		reqService.Id = uint64(id)
		var respService respPublisher

		err := NATS.RequestToNats("Publishers", "Web", "GetPublisher", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		name := respService.Name
		desc := respService.Description
		return c.Render(id, name, desc)
	}
}
func (c Publishers) Create(Name string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		var reqService requestCreatePublisher
		reqService.Name = Name
		reqService.Description = Description
		var respService respRemovePublisher

		err := NATS.RequestToNats("Publishers", "Web", "CreatePublisher", &reqService, &respService)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Login.Login)
		}

		return c.Redirect(Publishers.Publishers)
	}
	return c.Render()

}
