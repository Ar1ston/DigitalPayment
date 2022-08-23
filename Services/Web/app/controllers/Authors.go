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
type respAuthor struct {
	Id          uint64 `json:"id"`
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
	Id uint64 `json:"id"`
}
type requestRemoveAuthor struct {
	Id uint64 `json:"id"`
}
type requestChangeAuthor struct {
	Id          uint64 `json:"id"`
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
	req.To = "Authors"
	req.From = "Web"
	req.RequestName = "GetAuthors"
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
	var respService respAuthors
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

	var auth []author
	auth = respService.Authors

	return c.Render(auth)
}
func (c Authors) Author(id int) revel.Result {

	//конфиг
	var nats NATS.ConnectNATS
	nats.Host = "localhost"
	nats.Port = "4222"

	ConnNats, err := nats.ConnectToNATS()
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос в натц")

	//кодирование сообщения в натц

	var reqNats requestGetAuthor
	reqNats.Id = uint64(id)

	var buff4 bytes.Buffer
	enc := gob.NewEncoder(&buff4)
	err = enc.Encode(reqNats)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	//запрос в натц
	var req NATS.RequestNats
	req.Msg = buff4.Bytes()
	req.To = "Authors"
	req.From = "Web"
	req.RequestName = "GetAuthor"
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
	var respService respAuthor
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

	firstName := respService.FirstName
	lastName := respService.LastName
	desc := respService.Description

	return c.Render(id, firstName, lastName, desc)
}
func (c Authors) Remove(id int) revel.Result {

	//конфиг
	var nats NATS.ConnectNATS
	nats.Host = "localhost"
	nats.Port = "4222"

	ConnNats, err := nats.ConnectToNATS()
	if err != nil {
		return c.Redirect(Login.Login)
	}

	fmt.Println("Запрос в натц")

	//кодирование сообщения в натц

	var reqNats requestRemoveAuthor
	reqNats.Id = uint64(id)

	var buff4 bytes.Buffer
	enc := gob.NewEncoder(&buff4)
	err = enc.Encode(reqNats)
	if err != nil {
		return c.Redirect(Login.Login)
	}

	//запрос в натц
	var req NATS.RequestNats
	req.Msg = buff4.Bytes()
	req.To = "Authors"
	req.From = "Web"
	req.RequestName = "RemoveAuthor"
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
	var respService respRemoveAuthor
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

	return c.Redirect(Authors.Authors)
}
func (c Authors) Change(id int, FirstName string, LastName string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		//конфиг
		var nats NATS.ConnectNATS
		nats.Host = "localhost"
		nats.Port = "4222"

		ConnNats, err := nats.ConnectToNATS()
		if err != nil {
			return c.Redirect(Login.Login)
		}

		fmt.Println("Запрос в натц")

		//кодирование сообщения в натц

		var reqNats requestChangeAuthor
		reqNats.Id = uint64(id)
		reqNats.First_name = FirstName
		reqNats.Last_name = LastName
		reqNats.Description = Description

		var buff4 bytes.Buffer
		enc := gob.NewEncoder(&buff4)
		err = enc.Encode(reqNats)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		//запрос в натц
		var req NATS.RequestNats
		req.Msg = buff4.Bytes()
		req.To = "Authors"
		req.From = "Web"
		req.RequestName = "ChangeAuthor"
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
		var respService respRemoveAuthor
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

		return c.Redirect(Authors.Authors)
	} else {

		//конфиг
		var nats NATS.ConnectNATS
		nats.Host = "localhost"
		nats.Port = "4222"

		ConnNats, err := nats.ConnectToNATS()
		if err != nil {
			return c.Redirect(Login.Login)
		}

		fmt.Println("Запрос в натц")

		//кодирование сообщения в натц

		var reqNats requestGetAuthor
		reqNats.Id = uint64(id)

		var buff4 bytes.Buffer
		enc := gob.NewEncoder(&buff4)
		err = enc.Encode(reqNats)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		//запрос в натц
		var req NATS.RequestNats
		req.Msg = buff4.Bytes()
		req.To = "Authors"
		req.From = "Web"
		req.RequestName = "GetAuthor"
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
		var respService respAuthor
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

		firstName := respService.FirstName
		lastName := respService.LastName
		desc := respService.Description

		return c.Render(id, firstName, lastName, desc)
	}
}
func (c Authors) Create(FirstName string, LastName string, Description string) revel.Result {
	if c.Request.Method == "POST" {

		//конфиг
		var nats NATS.ConnectNATS
		nats.Host = "localhost"
		nats.Port = "4222"

		ConnNats, err := nats.ConnectToNATS()
		if err != nil {
			return c.Redirect(Login.Login)
		}

		fmt.Println("Запрос в натц")

		//кодирование сообщения в натц

		var reqNats requestCreateAuthor
		reqNats.First_name = FirstName
		reqNats.Last_name = LastName
		reqNats.Description = Description

		var buff4 bytes.Buffer
		enc := gob.NewEncoder(&buff4)
		err = enc.Encode(reqNats)
		if err != nil {
			return c.Redirect(Login.Login)
		}

		//запрос в натц
		var req NATS.RequestNats
		req.Msg = buff4.Bytes()
		req.To = "Authors"
		req.From = "Web"
		req.RequestName = "CreateAuthor"
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
		var respService respRemoveAuthor
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

		return c.Redirect(Authors.Authors)
	}
	return c.Render()
}
