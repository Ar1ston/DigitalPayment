package Requests

import (
	"DigitalPayment/Services/Books/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "ChangeBook"
	register_requests.Register(method, (*RequestChangeBook)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestChangeBook struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeBook) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле Id в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestChangeBook: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestChangeBook) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseChangeBook{}

	updates := map[string]interface{}{}

	if request.Name != "" {
		updates["Name"] = request.Name
	}
	if request.Genre != "" {
		updates["Genre"] = request.Genre
	}
	if request.Author != 0 {
		updates["Author"] = request.Author
	}
	if request.Publisher != 0 {
		updates["Publisher"] = request.Publisher
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	books, err := db_local.ChangeBookById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(books.Id)
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
