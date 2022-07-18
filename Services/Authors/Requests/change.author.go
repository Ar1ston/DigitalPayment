package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/Services/Authors/lib/reflect_local"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "ChangeAuthor"
	reflect_local.Register(method, (*RequestChangeAuthor)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestChangeAuthor struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeAuthor) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	if request.First_name == "" {
		err = fmt.Errorf("%s", "Неверное поле First_name в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	if request.Last_name == "" {
		err = fmt.Errorf("%s", "Неверное поле Last_name в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	if request.Description == "" {
		err = fmt.Errorf("%s", "Неверное поле Description в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestChangeAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseChangeAuthor{}

	updates := map[string]interface{}{}

	if request.First_name != "" {
		updates["FirstName"] = request.First_name
	}
	if request.Last_name != "" {
		updates["LastName"] = request.Last_name
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	author, err := db_local.ChangeAuthorById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(author.Id)
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
