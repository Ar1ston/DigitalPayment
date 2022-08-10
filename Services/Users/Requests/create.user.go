package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
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
	var rplBytes = bytes.NewBuffer(decReq)
	dec := gob.NewDecoder(rplBytes)
	err := dec.Decode(request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreateUser) Validation() *error {
	var err error
	if request.Login == "" {
		err = fmt.Errorf("%s", "Неверное поле Login в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCreateUser: %s\n", err.Error())
		return &err
	}
	if request.Password == "" {
		err = fmt.Errorf("%s", "Неверное поле Password в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCreateUser: %s\n", err.Error())
		return &err
	}
	return nil
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

	var rplBytes bytes.Buffer
	enc := gob.NewEncoder(&rplBytes)

	err = enc.Encode(rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes.Bytes(), nil
}
