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
	method := "GetBooks"
	register_requests.Register(method, (*RequestGetBooks)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetBooks struct{}
type Book struct {
	Id          int64
	Name        string
	Genre       string
	Author      int64
	Publisher   int64
	AddedUser   int64
	AddedTime   time.Time
	Description string
}
type ResponseGetBooks struct {
	Books []Book
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestGetBooks) Validation() *error {

	return nil
}
func (request *RequestGetBooks) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetBooks{}

	books, err := db_local.FindBooks(db_local.DB_LOCAL, map[string]interface{}{})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range books.Books {
			rpl.Books = append(rpl.Books, Book{
				Id:          v.Id,
				Name:        v.Name,
				Genre:       v.Genre,
				Author:      v.Author,
				Publisher:   v.Publisher,
				AddedUser:   v.AddedUser,
				AddedTime:   v.AddedTime,
				Description: v.Description,
			})
		}
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
