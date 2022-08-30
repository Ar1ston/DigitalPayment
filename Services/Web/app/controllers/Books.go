package controllers

import (
	"Web/conf/NATS"
	"fmt"
	"github.com/revel/revel"
	"time"
)

type Books struct {
	*revel.Controller
}
type book struct {
	Id     int64
	Name   string
	Genre  string
	Author int64
}
type respBooks struct {
	Books []book
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

type requestGetBook struct {
	Id uint64 `json:"Id"`
}
type respGetBook struct {
	Id          uint64    `json:"Id"`
	Name        string    `json:"Name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
	Errno       uint64    `json:"errno"`
	Error       string    `json:"error,omitempty"`
}

type requestCreateBook struct {
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Added_User  uint64 `json:"addedUser,omitempty"`
	Description string `json:"description,omitempty"`
}
type respCreateBook struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

type BookPublishers struct {
	Id   int
	Name string
}
type BookUsers struct {
	Id    int
	Login string
}
type BookAuthors struct {
	Id   int
	Name string
}

func (c Books) Books() revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	var respService respBooks
	err := NATS.RequestToNats("Books", "Web", "GetBooks", []byte(""), &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	var bks []book
	bks = respService.Books
	return c.Render(bks)
}
func (c Books) Book(id int) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	var reqService requestGetBook
	reqService.Id = uint64(id)
	var respService respGetBook

	err := NATS.RequestToNats("Books", "Web", "GetBook", &reqService, &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	var name = respService.Name
	var genre = respService.Genre
	var author = respService.Author
	var publisher = respService.Publisher
	var added_User = respService.Added_User
	var added_Time = respService.Added_Time.Format("02.01.2006")
	var description = respService.Description

	return c.Render(name, genre, author, publisher, added_User, added_Time, description)
}
func (c Books) Create(publishers []BookPublishers, users []BookUsers, authors []BookAuthors, Name string, Genre string, author int, publisher int, user int, Description string) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	if c.Session["level"] == "1" {
		return c.Redirect(Error.Error, 409, "No access")
	}

	if c.Request.Method == "GET" {

		//Запрашиваем авторов
		var respServiceAuthor respAuthors
		err := NATS.RequestToNats("Authors", "Web", "GetAuthors", []byte(""), &respServiceAuthor)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}
		if respServiceAuthor.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServiceAuthor.Errno, respServiceAuthor.Error)
			return c.Redirect(Error.Error, int(respServiceAuthor.Errno), respServiceAuthor.Error)
		}
		var authors []BookAuthors
		for _, v := range respServiceAuthor.Authors {
			authors = append(authors, BookAuthors{
				Id:   int(v.Id),
				Name: v.FirstName + " " + v.LastName,
			})
		}

		//Запрашиваем пользователей
		var respServiceUser respGetUsers
		err = NATS.RequestToNats("Users", "Web", "GetUsers", []byte(""), &respServiceUser)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}
		if respServiceUser.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServiceUser.Errno, respServiceUser.Error)
			return c.Redirect(Error.Error, int(respServiceUser.Errno), respServiceUser.Error)
		}
		var users []BookUsers
		for _, v := range respServiceUser.Users {
			users = append(users, BookUsers{
				Id:    int(v.Id),
				Login: v.Login,
			})
		}

		//Запрашиваем публикаторов
		var respServicePublisher respPublishers
		err = NATS.RequestToNats("Publishers", "Web", "GetPublishers", []byte(""), &respServicePublisher)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}
		if respServicePublisher.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServicePublisher.Errno, respServicePublisher.Error)
			return c.Redirect(Error.Error, int(respServicePublisher.Errno), respServicePublisher.Error)
		}
		var publishers []BookPublishers
		for _, v := range respServicePublisher.Publishers {
			publishers = append(publishers, BookPublishers{
				Id:   int(v.Id),
				Name: v.Name,
			})
		}

		c.Render(publishers, users, authors)
	}
	if c.Request.Method == "POST" {
		fmt.Printf("Name: %s, Genre: %s, Author: %d, Publisher: %d, User: %d, Desc:%s\n", Name, Genre, author, publisher, user, Description)

		var reqService requestCreateBook
		reqService.Name = Name
		reqService.Genre = Genre
		reqService.Author = uint64(author)
		reqService.Publisher = uint64(publisher)
		reqService.Added_User = uint64(user)
		reqService.Description = Description
		var respService respCreateBook

		err := NATS.RequestToNats("Books", "Web", "CreateBook", &reqService, &respService)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
		}

		return c.Redirect(Books.Books)

	}
	return c.Render()

}
