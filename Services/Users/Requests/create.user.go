package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
	"fmt"
)

func init() {
	method := "CreateUser"
	register_requests.Register(method, (*RequestCreateUser)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCreateUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type ResponseCreateUser struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateUser) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreateUser) Validation() []byte {
	isError := false
	rpl := ResponseCreateUser{}
	if request.Login == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Login field in request"
		fmt.Printf("ERROR VALIDATION: %s\n", rpl.Error)
	}
	if request.Password == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Password field in request"
		fmt.Printf("ERROR VALIDATION: %s\n", rpl.Error)
	}
	if isError == false {
		return nil
	} else {
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestCreateUser) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCreateUser{}

	req := db_local.User{
		Login:    request.Login,
		Password: request.Password,
		Name:     request.Name,
	}

	user, err := db_local.CreateUser(db_local.DB_LOCAL, &req)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(user.Id)
		rpl.Errno = 0
	}
	fmt.Printf("RESPONSE: %+v\n", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
