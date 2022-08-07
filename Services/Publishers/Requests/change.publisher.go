package Requests

import (
	"DigitalPayment/Services/Publishers/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func Hello() {}
func init() {
	method := "ChangePublisher"
	register_requests.Register(method, (*RequestChangePublisher)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestChangePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangePublisher) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле Id в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestChangePublisher: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestChangePublisher) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseChangePublisher{}

	updates := map[string]interface{}{}

	if request.Name != "" {
		updates["Name"] = request.Name
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	publisher, err := db_local.ChangePublisherById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(publisher.Id)
		rpl.Errno = 0
	}
	fmt.Printf("RESPONSE: %+v\n", rpl)

	var rplBytes bytes.Buffer
	enc := gob.NewEncoder(&rplBytes)

	err = enc.Encode(rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes.Bytes(), nil
}
