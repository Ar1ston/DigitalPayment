package Requests

import (
	"DigitalPayment/Services/Publishers/lib/db_local"
	"DigitalPayment/Services/Publishers/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "CreatePublisher"
	register_requests.Register(method, (*RequestCreatePublisher)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreatePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreatePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreatePublisher) Validation() *error {
	var err error
	if request.Name == "" {
		err = fmt.Errorf("%s", "Неверное поле Name в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCreatePublisher: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestCreatePublisher) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCreatePublisher{}

	req := db_local.Publisher{
		Name:        request.Name,
		Description: request.Description,
	}

	publisher, err := db_local.CreatePublisher(db_local.DB_LOCAL, &req)

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
