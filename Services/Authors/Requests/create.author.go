package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
	"fmt"
)

func init() {
	method := "CreateAuthor"
	register_requests.Register(method, (*RequestCreateAuthor)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreateAuthor struct {
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreateAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateAuthor) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreateAuthor) Validation() []byte {
	isError := false
	rpl := ResponseCreateAuthor{}
	if request.First_name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation FirstName field in request"
		fmt.Printf("ERROR VALIDATION RequestGetAuthor: %s\n", rpl.Error)
	}
	if request.Last_name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation LastName field in request"
		fmt.Printf("ERROR VALIDATION RequestGetAuthor: %s\n", rpl.Error)
	}
	if isError == false {
		return nil
	} else {
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt

	}
}
func (request *RequestCreateAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCreateAuthor{}

	author, err := db_local.CreateAuthor(db_local.DB_LOCAL, request.First_name, request.Last_name, request.Description)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(author.Id)
		rpl.Errno = 0
	}
	fmt.Printf("RESPONSE: %+v\n", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
