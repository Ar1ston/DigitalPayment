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
type respRemovePublisher struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}
type requestGetPublisher struct {
	Id uint64 `json:"id"`
}
type requestRemovePublisher struct {
	Id uint64 `json:"id"`
}
type requestChangePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type requestCreatePublisher struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
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

	name := respService.Name
	desc := respService.Description
	return c.Render(id, name, desc)
}
func (c Publishers) Remove(id int) revel.Result {
	fmt.Printf("%d\n", id)
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

	var reqNats requestRemovePublisher
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
	req.RequestName = "RemovePublisher"
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

	return c.Redirect(Publishers.Publishers)
}
func (c Publishers) Change(id int, Name string, Description string) revel.Result {
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

		var reqNats requestChangePublisher
		reqNats.Id = uint64(id)
		reqNats.Name = Name
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
		req.To = "Publishers"
		req.From = "Web"
		req.RequestName = "ChangePublisher"
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
		var respService respRemovePublisher
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

		return c.Redirect(Publishers.Publishers)
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

		var reqNats requestGetPublisher
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

		name := respService.Name
		desc := respService.Description

		return c.Render(id, name, desc)
	}
}
func (c Publishers) Create(Name string, Description string) revel.Result {
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

		var reqNats requestCreatePublisher
		reqNats.Name = Name
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
		req.To = "Publishers"
		req.From = "Web"
		req.RequestName = "CreatePublisher"
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
		var respService respRemovePublisher
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

		return c.Redirect(Publishers.Publishers)
	}
	return c.Render()

}
