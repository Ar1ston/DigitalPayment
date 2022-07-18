package Requests

import (
	"DigitalPayment/Services/Books/lib/reflect_local"
	"fmt"
)

func init() {
	method := "ChangeBook"
	reflect_local.Register(method, (*RequestChangeBook)(nil))
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

	return nil
}
func (request *RequestChangeBook) Execute() ([]byte, *error) {

	return nil, nil
}
