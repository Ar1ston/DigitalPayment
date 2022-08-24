package controllers

import (
	"Web/conf/NATS"
	"bytes"
	"encoding/gob"
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
type respBook struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
	Errno       uint64    `json:"errno"`
	Error       string    `json:"error,omitempty"`
}

type BookPublishers struct {
	id   int
	name string
}
type BookUsers struct {
	id    int
	login string
}
type BookAuthors struct {
	id   int
	name string
}

func (c Books) Books() revel.Result {

	//конфиг
	var nats NATS.ConnectNATS
	nats.Host = "localhost"
	nats.Port = "4222"

	ConnNats, err := nats.ConnectToNATS()
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос в натц")
	//запрос в натц
	var req NATS.RequestNats
	req.Msg = []byte("")
	req.To = "Books"
	req.From = "Web"
	req.RequestName = "GetBooks"
	rpl, err := req.SendRequestToNats(ConnNats)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос сделан")

	fmt.Println("Декодирование 1")
	//декодирование сообщения натц
	resp := NATS.RequestNats{}
	var rplBytes = bytes.NewBuffer(rpl)
	dec := gob.NewDecoder(rplBytes)
	err = dec.Decode(&resp)
	if err != nil {
		return c.Redirect(Login.Login)
	}
	fmt.Println("Декодирование 1. Конец.")

	fmt.Println("Декодирование 2")
	//декодирование сообщения сервиса
	fmt.Printf("FROM: %s TO: %s ReqName: %s\n", resp.From, resp.To, resp.RequestName)
	var respService respBooks
	var respServBytes = bytes.NewBuffer(resp.Msg)
	dec = gob.NewDecoder(respServBytes)
	err = dec.Decode(&respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}
	fmt.Println("Декодирование 2. Конец.")

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
	}

	var bks []book
	bks = respService.Books
	return c.Render(bks)
}
func (c Books) Book(id int) revel.Result {
	//конфиг
	var nats NATS.ConnectNATS
	nats.Host = "localhost"
	nats.Port = "4222"

	ConnNats, err := nats.ConnectToNATS()
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос в натц")
	//запрос в натц

	//кодирование сообщения в натц

	var reqNats requestGetPublisher
	reqNats.Id = uint64(id)

	var buff4 bytes.Buffer
	enc := gob.NewEncoder(&buff4)
	err = enc.Encode(reqNats)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println(string(buff4.Bytes()))
	var req NATS.RequestNats
	req.Msg = buff4.Bytes()
	req.To = "Books"
	req.From = "Web"
	req.RequestName = "GetBook"
	rpl, err := req.SendRequestToNats(ConnNats)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос сделан")

	fmt.Println("Декодирование 1")
	//декодирование сообщения натц
	resp := NATS.RequestNats{}
	var rplBytes = bytes.NewBuffer(rpl)
	dec := gob.NewDecoder(rplBytes)
	err = dec.Decode(&resp)
	if err != nil {
		return c.Redirect(Login.Login)
	}
	fmt.Println("Декодирование 1. Конец.")

	fmt.Println("Декодирование 2")
	//декодирование сообщения сервиса
	fmt.Printf("FROM: %s TO: %s ReqName: %s\n", resp.From, resp.To, resp.RequestName)
	var respService respBook
	var respServBytes = bytes.NewBuffer(resp.Msg)
	dec = gob.NewDecoder(respServBytes)
	err = dec.Decode(&respService)
	if err != nil {
		return c.Redirect(Login.Login)
	}
	fmt.Println("Декодирование 2. Конец.")

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(Login.Login)
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
func (c Books) Create(publishers []BookPublishers, users []BookUsers, authors []BookAuthors) revel.Result {
	//if c.Request.Method == "POST" {
	//
	//	//конфиг
	//	var nats NATS.ConnectNATS
	//	nats.Host = "localhost"
	//	nats.Port = "4222"
	//
	//	ConnNats, err := nats.ConnectToNATS()
	//	if err != nil {
	//		return c.Redirect(Login.Login)
	//	}
	//
	//	fmt.Println("Запрос в натц")
	//
	//	//кодирование сообщения в натц
	//
	//	var reqNats requestCreatePublisher
	//	reqNats.Name = Name
	//	reqNats.Description = Description
	//
	//	var buff4 bytes.Buffer
	//	enc := gob.NewEncoder(&buff4)
	//	err = enc.Encode(reqNats)
	//	if err != nil {
	//		return c.Redirect(Login.Login)
	//	}
	//
	//	//запрос в натц
	//	var req NATS.RequestNats
	//	req.Msg = buff4.Bytes()
	//	req.To = "Publishers"
	//	req.From = "Web"
	//	req.RequestName = "CreatePublisher"
	//	rpl, err := req.SendRequestToNats(ConnNats)
	//	if err != nil {
	//		return c.Redirect(Login.Login)
	//	}
	//
	//	fmt.Println("Запрос сделан")
	//
	//	fmt.Println("Декодирование 1")
	//	//декодирование сообщения натц
	//	resp := NATS.RequestNats{}
	//	var rplBytes = bytes.NewBuffer(rpl)
	//	dec := gob.NewDecoder(rplBytes)
	//	err = dec.Decode(&resp)
	//	if err != nil {
	//		return c.Redirect(Login.Login)
	//	}
	//	fmt.Println("Декодирование 1. Конец.")
	//
	//	fmt.Println("Декодирование 2")
	//	//декодирование сообщения сервиса
	//	fmt.Printf("FROM: %s TO: %s ReqName: %s\n", resp.From, resp.To, resp.RequestName)
	//	var respService respRemovePublisher
	//	var respServBytes = bytes.NewBuffer(resp.Msg)
	//	dec = gob.NewDecoder(respServBytes)
	//	err = dec.Decode(&respService)
	//	if err != nil {
	//		return c.Redirect(Login.Login)
	//	}
	//	fmt.Println("Декодирование 2. Конец.")
	//
	//	if respService.Errno != 0 {
	//		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
	//		return c.Redirect(Login.Login)
	//	}
	//
	//	return c.Redirect(Publishers.Publishers)
	//}else{
	return c.Render()
	//}
}
