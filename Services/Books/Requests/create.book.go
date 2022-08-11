package Requests

import (
	"DigitalPayment/Services/Books/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

func init() {
	method := "CreateBook"
	register_requests.Register(method, (*RequestCreateBook)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreateBook struct {
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Added_User  uint64 `json:"addedUser,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreateBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateBook) Decode(decReq []byte) *error {
	var rplBytes = bytes.NewBuffer(decReq)
	dec := gob.NewDecoder(rplBytes)
	err := dec.Decode(request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreateBook) Validation() *error {
	var err error
	if request.Name == "" {
		err = fmt.Errorf("%s", "Неверное поле Name в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCreateBook: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestCreateBook) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCreateBook{}

	req := db_local.Book{
		Name:        request.Name,
		Genre:       request.Genre,
		Author:      int64(request.Author),
		Publisher:   int64(request.Publisher),
		AddedUser:   int64(request.Added_User),
		AddedTime:   time.Now(),
		Description: request.Description,
	}

	//TODO не факт что заработает
	book, err := db_local.CreateBook(db_local.DB_LOCAL, &req)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(book.Id)
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
