package Requests

import (
	"DigitalPayment/Services/Publishers/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
	"fmt"
)

func init() {
	method := "CreatePublisher"
	register_requests.Register(method, (*RequestCreatePublisher)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreatePublisher struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreatePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreatePublisher) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreatePublisher) Validation() []byte {
	isError := false
	rpl := ResponseCreatePublisher{}
	if request.Name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Name field in request"
		fmt.Printf("ERROR VALIDATION CreatePublisher: %s\n", rpl.Error)
	}
	if isError == false {
		return nil
	} else {
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
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

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
