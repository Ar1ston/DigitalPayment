package Requests

import (
	"DigitalPayment/Services/Books/lib/reflect_local"
	"fmt"
	"time"
)

func init() {
	method := "GetBooks"
	reflect_local.Register(method, (*RequestGetBooks)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetBooks struct{}
type Book struct {
	Id          int64 `json:"id"`
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

	return nil, nil
}
