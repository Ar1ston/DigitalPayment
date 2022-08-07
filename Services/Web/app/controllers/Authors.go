package controllers

import (
	"Web/conf/NATS"
	"bytes"
	"encoding/gob"
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

func (c Authors) Authors() revel.Result {

	//конфиг
	var nats NATS.ConnectNATS
	nats.Host = "localhost"
	nats.Port = "4222"

	ConnNats, err := nats.ConnectToNATS()
	if err != nil {
		c.Redirect(App.Index)
	}

	fmt.Println("Запрос в натц")
	//запрос в натц
	var req NATS.RequestNats
	req.Msg = []byte("")
	req.To = "Authors"
	req.From = "Web"
	req.RequestName = "GetAuthors"
	rpl, err := req.SendRequestToNats(ConnNats)
	if err != nil {
		c.Redirect(App.Index)
	}

	fmt.Println("Запрос сделан")

	fmt.Println("Декодирование 1")
	//декодирование сообщения натц
	resp := NATS.RequestNats{}
	var rplBytes = bytes.NewBuffer(rpl)
	dec := gob.NewDecoder(rplBytes)
	err = dec.Decode(&resp)
	if err != nil {
		c.Redirect(App.Index)
	}
	fmt.Println("Декодирование 1. Конец.")

	fmt.Println("Декодирование 2")
	//декодирование сообщения сервиса
	fmt.Printf("FROM: %s TO: %s ReqName: %s\n", resp.From, resp.To, resp.RequestName)
	var respService respAuthors
	var respServBytes = bytes.NewBuffer(resp.Msg)
	dec = gob.NewDecoder(respServBytes)
	err = dec.Decode(&respService)
	if err != nil {
		c.Redirect(App.Index)
	}
	fmt.Println("Декодирование 2. Конец.")

	if respService.Errno != 0 {
		fmt.Printf("ERROR SERVICE(code %d): %s", respService.Errno, respService.Error)
		return c.Redirect(App.Index)
	}

	var auth []author
	auth = respService.Authors
	//auth = append(auth, author{FirstName: "Name",
	//	LastName:    "LastName",
	//	Description: "Desc"})
	//auth = append(auth, author{FirstName: "Name2",
	//	LastName:    "LastName2",
	//	Description: "Desc2"})
	//auth = append(auth, author{FirstName: "Name3",
	//	LastName:    "LastName3",
	//	Description: "Desc3"})
	return c.Render(auth)
}
