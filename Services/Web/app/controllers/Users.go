package controllers

import (
	"Web/conf/NATS"
	"bytes"
	"encoding/gob"
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
type respUsers struct {
	Users []user
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}
type respUser struct {
	Id    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Level uint64 `json:"level"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (c Users) Users() revel.Result {

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
	req.To = "Users"
	req.From = "Web"
	req.RequestName = "GetUsers"
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
	var respService respUsers
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

	var usrs []user
	usrs = respService.Users
	//auth = append(auth, author{FirstName: "Name",
	//	LastName:    "LastName",
	//	Description: "Desc"})
	//auth = append(auth, author{FirstName: "Name2",
	//	LastName:    "LastName2",
	//	Description: "Desc2"})
	//auth = append(auth, author{FirstName: "Name3",
	//	LastName:    "LastName3",
	//	Description: "Desc3"})
	return c.Render(usrs)
}
func (c Users) User(id int) revel.Result {

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

	var reqNats user
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
	req.To = "Users"
	req.From = "Web"
	req.RequestName = "GetUser"
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
	var respService respUser
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

	Login := respService.Login
	Name := respService.Name
	Level := respService.Level

	return c.Render(Name, Login, Level)
}
