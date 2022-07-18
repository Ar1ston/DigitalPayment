package Requests

import (
	"DigitalPayment/Services/Books/lib/reflect_local"
	"fmt"
	"time"
)

func init() {
	method := "GetBook"
	reflect_local.Register(method, (*RequestGetBook)(nil))
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

	return nil
}
func (request *RequestGetBook) Execute() ([]byte, *error) {

	return nil, nil
}
