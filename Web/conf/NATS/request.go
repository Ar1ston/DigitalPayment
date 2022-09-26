package NATS

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type RequestNats struct {
	Msg         []byte
	To          string
	RequestName string
	From        string
}

func (req *RequestNats) SendRequestToNats(nats *nats.Conn) ([]byte, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(req)
	if err != nil {
		return nil, err
	}

	resp, err := nats.Request(req.To, buff.Bytes(), time.Second)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
func toBytes(message interface{}) []byte {

	js, err := json.Marshal(message)
	if err != nil {
		log.Printf("Unable to marshal: %+v: %s", message, err)
		return nil
	}

	return js
}
func RequestToNats(To string, From string, RequestName string, Request interface{}, Response interface{}) error {
	fmt.Println("Начало отправки запроса в сервис через натс")

	//конфиг
	var connectNATS ConnectNATS
	connectNATS.Host = "192.168.20.141"
	connectNATS.Port = "4222"

	ConnNats, err := connectNATS.ConnectToNATS()
	if err != nil {
		return err
	}

	fmt.Println("Запрос в натц")

	//кодирование сообщения в натц
	fmt.Println("Кодирование запроса сервиса")

	var buff4 bytes.Buffer
	enc := gob.NewEncoder(&buff4)
	err = enc.Encode(Request)
	if err != nil {
		return err
	}
	fmt.Println("Кодирование запроса сервиса. Конец")

	fmt.Println("Запрос в натц")

	//запрос в натц
	var req RequestNats
	req.Msg = buff4.Bytes()
	req.To = To
	req.From = From
	req.RequestName = RequestName
	rpl, err := req.SendRequestToNats(ConnNats)
	if err != nil {
		return err
	}

	fmt.Println("Запрос сделан")

	fmt.Println("Декодирование 1")
	//декодирование сообщения натц
	resp := RequestNats{}
	var rplBytes = bytes.NewBuffer(rpl)
	dec := gob.NewDecoder(rplBytes)
	err = dec.Decode(&resp)
	if err != nil {
		return err
	}
	fmt.Println("Декодирование 1. Конец.")

	fmt.Println("Декодирование 2")
	//декодирование сообщения сервиса
	fmt.Printf("FROM: %s TO: %s ReqName: %s\n", resp.From, resp.To, resp.RequestName)
	var respServBytes = bytes.NewBuffer(resp.Msg)
	dec = gob.NewDecoder(respServBytes)
	err = dec.Decode(Response)
	if err != nil {
		return err
	}
	fmt.Println("Декодирование 2. Конец.")

	return nil
}
