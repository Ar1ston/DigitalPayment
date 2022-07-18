package Requests

import (
	"DigitalPayment/Services/Books/lib/reflect_local"
	"fmt"
	"time"
)

func init() {
	method := "CreateBook"
	reflect_local.Register(method, (*RequestCreateBook)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreateBook struct {
	Name        string    `json:"name,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Author      uint64    `json:"author,omitempty"`
	Publisher   uint64    `json:"publisher,omitempty"`
	Added_User  uint64    `json:"addedUser,omitempty"`
	Added_Time  time.Time `json:"addedTime,omitempty"`
	Description string    `json:"description,omitempty"`
}
type ResponseCreateBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateBook) Validation() *error {

	return nil
}
func (request *RequestCreateBook) Execute() ([]byte, *error) {

	return nil, nil
}
