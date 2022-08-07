package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "GetUsers"
	register_requests.Register(method, (*RequestGetUsers)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetUsers struct{}

type user struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Level uint64 `json:"level"`
}
type ResponseGetUsers struct {
	Users []user
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestGetUsers) Validation() *error {
	return nil
}
func (request *RequestGetUsers) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetUsers{}

	users, err := db_local.FindUsers(db_local.DB_LOCAL, map[string]interface{}{})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range users.Users {
			rpl.Users = append(rpl.Users, user{
				Id:    uint64(v.Id),
				Name:  v.Name,
				Level: uint64(v.Level),
			})
		}
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
