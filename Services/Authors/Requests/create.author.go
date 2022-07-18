package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/Services/Authors/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "CreateAuthor"
	register_requests.Register(method, (*RequestCreateAuthor)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreateAuthor struct {
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreateAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateAuthor) Validation() *error {
	var err error
	if request.First_name == "" {
		err = fmt.Errorf("%s", "Неверное поле FirstName в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	if request.Last_name == "" {
		err = fmt.Errorf("%s", "Неверное поле LastName в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestCreateAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCreateAuthor{}

	author, err := db_local.CreateAuthor(db_local.DB_LOCAL, request.First_name, request.Last_name, request.Description)

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
