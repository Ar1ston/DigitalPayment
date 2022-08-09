package controllers

import (
	"Web/conf/NATS"
	"bytes"
	"encoding/gob"
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
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}
type requestGetPublisher struct {
	Id uint64 `json:"id"`
}

func (c Publishers) Publishers() revel.Result {

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
	req.To = "Publishers"
	req.From = "Web"
	req.RequestName = "GetPublishers"
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
	var respService respPublishers
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

	var pubs []publisher
	pubs = respService.Publishers
	//auth = append(auth, author{FirstName: "Name",
	//	LastName:    "LastName",
	//	Description: "Desc"})
	//auth = append(auth, author{FirstName: "Name2",
	//	LastName:    "LastName2",
	//	Description: "Desc2"})
	//auth = append(auth, author{FirstName: "Name3",
	//	LastName:    "LastName3",
	//	Description: "Desc3"})
	return c.Render(pubs)
}
func (c Publishers) Publisher(id int) revel.Result {
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
	req.To = "Publishers"
	req.From = "Web"
	req.RequestName = "GetPublisher"
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
	var respService respPublisher
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

	//var pub = struct {
	//	Id          uint64
	//	Name        string
	//	Description string
	//}{
	//	Name:        respService.Name,
	//	Description: respService.Description,
	//}
	name := respService.Name
	desc := respService.Description
	return c.Render(name, desc)
}
