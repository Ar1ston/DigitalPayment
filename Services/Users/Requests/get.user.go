package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "GetUser"
	register_requests.Register(method, (*RequestGetUser)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetUser struct {
	Login string `json:"login"`
}
type ResponseGetUser struct {
	Id       uint64 `json:"id"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
	Errno    uint64 `json:"errno"`
	Error    string `json:"error,omitempty"`
}

func (request *RequestGetUser) Decode(decReq []byte) *error {
	var rplBytes = bytes.NewBuffer(decReq)
	dec := gob.NewDecoder(rplBytes)
	err := dec.Decode(request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestGetUser) Validation() *error {
	var err error
	if request.Login == "" {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetUser: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestGetUser) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetUser{}

	user, err := db_local.FindUser(db_local.DB_LOCAL, map[string]interface{}{
		"login": request.Login,
	})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(user.Id)
		rpl.Password = user.Password
		rpl.Name = user.Name
		rpl.Level = uint64(user.Level)
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
