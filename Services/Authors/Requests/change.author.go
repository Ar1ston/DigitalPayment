package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
	"fmt"
)

func Hello() {}
func init() {
	method := "ChangeAuthor"
	register_requests.Register(method, (*RequestChangeAuthor)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestChangeAuthor struct {
	Id          uint64 `json:"id"`
	First_name  string `json:"firstName,omitempty"`
	Last_name   string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeAuthor struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeAuthor) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}

func (request *RequestChangeAuthor) Validation() []byte {
	isError := false
	rpl := ResponseChangeAuthor{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
		fmt.Printf("ERROR VALIDATION ChangeAuthor: %s\n", rpl.Error)
	}
	if request.First_name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation First_name field in request"
		fmt.Printf("ERROR VALIDATION ChangeAuthor: %s\n", rpl.Error)
	}
	if request.Last_name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Last_name field in request"
		fmt.Printf("ERROR VALIDATION ChangeAuthor: %s\n", rpl.Error)
	}
	if request.Description == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Description field in request"
		fmt.Printf("ERROR VALIDATION ChangeAuthor: %s\n", rpl.Error)
	}
	if isError == false {
		return nil
	} else {
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt

	}
}
func (request *RequestChangeAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseChangeAuthor{}

	updates := map[string]interface{}{}

	if request.First_name != "" {
		updates["FirstName"] = request.First_name
	}
	if request.Last_name != "" {
		updates["LastName"] = request.Last_name
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	author, err := db_local.ChangeAuthorById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

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
