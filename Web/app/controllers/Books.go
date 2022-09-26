package controllers

import (
	"Web/conf/NATS"
	"fmt"
	"github.com/revel/revel"
	"strconv"
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
type requestRemoveBook struct {
	Id uint64 `json:"id"`
}
type respRemoveBook struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}
type requestChangeBook struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Description string `json:"description,omitempty"`
}
type respChangeBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
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

	//if respService.Errno != 0 {
	//	fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
	//	return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	//}

	var bks []book
	bks = respService.Books
	return c.Render(bks)
}
func (c Books) Book(id int) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}

	var reqServiceBook requestGetBook
	reqServiceBook.Id = uint64(id)
	var respServiceBook respGetBook

	err := NATS.RequestToNats("Books", "Web", "GetBook", &reqServiceBook, &respServiceBook)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respServiceBook.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respServiceBook.Errno, respServiceBook.Error)
		return c.Redirect(Error.Error, int(respServiceBook.Errno), respServiceBook.Error)
	}

	var reqServiceAuthor requestGetAuthor
	reqServiceAuthor.Id = respServiceBook.Author
	var respServiceAuthor respAuthor

	err = NATS.RequestToNats("Authors", "Web", "GetAuthor", &reqServiceAuthor, &respServiceAuthor)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respServiceAuthor.Errno == 404 {
		respServiceAuthor.FirstName = "Неизвестно"
		respServiceAuthor.LastName = ""
	} else {
		if respServiceAuthor.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServiceAuthor.Errno, respServiceAuthor.Error)
			return c.Redirect(Error.Error, int(respServiceAuthor.Errno), respServiceAuthor.Error)
		}
	}

	var reqServicePublisher requestGetPublisher
	reqServicePublisher.Id = respServiceBook.Publisher
	var respServicePublisher respPublisher

	err = NATS.RequestToNats("Publishers", "Web", "GetPublisher", &reqServicePublisher, &respServicePublisher)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respServicePublisher.Errno == 404 {
		respServicePublisher.Name = "Неизвестно"
	} else {
		if respServicePublisher.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServicePublisher.Errno, respServicePublisher.Error)
			return c.Redirect(Error.Error, int(respServicePublisher.Errno), respServicePublisher.Error)
		}
	}

	var reqServiceUser requestGetUser
	reqServiceUser.Id = respServiceBook.Added_User
	var respServiceUser respGetUser

	err = NATS.RequestToNats("Users", "Web", "GetUser", &reqServiceUser, &respServiceUser)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respServiceUser.Errno == 404 {
		respServiceUser.Login = "Неизвестно"
	} else {
		if respServiceUser.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respServiceUser.Errno, respServiceUser.Error)
			return c.Redirect(Error.Error, int(respServiceUser.Errno), respServiceUser.Error)
		}
	}

	var name = respServiceBook.Name
	var genre = respServiceBook.Genre
	var id_author = respServiceBook.Author
	var id_publisher = respServiceBook.Publisher
	var author = respServiceAuthor.FirstName + " " + respServiceAuthor.LastName
	var publisher = respServicePublisher.Name
	var added_User = respServiceUser.Login
	var added_Time = respServiceBook.Added_Time.Format("02.01.2006")
	var description = respServiceBook.Description
	var user_id = respServiceBook.Added_User

	return c.Render(id, name, genre, author, publisher, added_User, added_Time, description, id_author, id_publisher, user_id)
}
func (c Books) Create(publishers []BookPublishers, users []BookUsers, authors []BookAuthors, Name string, Genre string, author int, publisher int, Description string) revel.Result {

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
		fmt.Printf("Name: %s, Genre: %s, Author: %d, Publisher: %d, User: %d, Desc:%s\n", Name, Genre, author, publisher, Description)

		var reqService requestCreateBook
		reqService.Name = Name
		reqService.Genre = Genre
		reqService.Author = uint64(author)
		reqService.Publisher = uint64(publisher)
		id, _ := strconv.Atoi(c.Session["id"].(string))
		reqService.Added_User = uint64(id)
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
func (c Books) Remove(id int) revel.Result {
	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}
	if c.Session["level"] != "3" {
		return c.Redirect(Error.Error, 409, "No access")
	}

	var reqService requestRemoveBook
	reqService.Id = uint64(id)

	var respService respRemoveBook

	err := NATS.RequestToNats("Books", "Web", "RemoveBook", &reqService, &respService)
	if err != nil {
		return c.Redirect(Error.Error, 500, "Error server")
	}

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
	}

	return c.Redirect(Books.Books)
}
func (c Books) Change(id int, publishers []BookPublishers, users []BookUsers, authors []BookAuthors, Name string, Genre string, author int, publisher int, Description string) revel.Result {

	if c.Session["login"] == nil {
		return c.Redirect(Login.Login)
	}
	if c.Session["level"] != "3" {
		return c.Redirect(Error.Error, 409, "No access")
	}
	if c.Request.Method == "POST" {

		var reqService requestChangeBook
		reqService.Id = uint64(id)
		reqService.Name = Name
		reqService.Genre = Genre
		reqService.Author = uint64(author)
		reqService.Publisher = uint64(publisher)
		reqService.Description = Description

		var respService respChangeBook

		err := NATS.RequestToNats("Books", "Web", "ChangeBook", &reqService, &respService)
		if err != nil {
			return c.Redirect(Error.Error, 500, "Error server")
		}

		if respService.Errno != 0 {
			fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
			return c.Redirect(Error.Error, int(respService.Errno), respService.Error)
		}

		return c.Redirect("/Book?id=%d", id)
	} else {

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

		//Запрашиваем авторов
		var respServiceAuthor respAuthors
		err = NATS.RequestToNats("Authors", "Web", "GetAuthors", []byte(""), &respServiceAuthor)
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

		var name = respService.Name
		var genre = respService.Genre
		var Selected_author = int(respService.Author)
		var selected_publisher = int(respService.Publisher)
		var description = respService.Description

		return c.Render(id, name, genre, authors, Selected_author, selected_publisher, publishers, description)
	}
}
