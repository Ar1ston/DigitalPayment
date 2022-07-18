package Requests

import (
	"DigitalPayment/Services/Books/lib/db_local"
	"DigitalPayment/Services/Books/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

func init() {
	method := "GetBook"
	register_requests.Register(method, (*RequestGetBook)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetBook struct {
	Id   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}
type ResponseGetBook struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
	Errno       uint64    `json:"errno"`
	Error       string    `json:"error,omitempty"`
}

func (request *RequestGetBook) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetBook: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestGetBook) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetBook{}

	book, err := db_local.FindBookById(db_local.DB_LOCAL, map[string]interface{}{
		"id": request.Id,
	})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(book.Id)
		rpl.Name = book.Name
		rpl.Genre = book.Genre
		rpl.Author = uint64(book.Author)
		rpl.Publisher = uint64(book.Publisher)
		rpl.Added_User = uint64(book.AddedUser)
		rpl.Added_Time = book.AddedTime
		rpl.Description = book.Description
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
