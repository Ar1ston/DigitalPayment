package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "RemoveUser"
	register_requests.Register(method, (*RequestRemoveUser)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestRemoveUser struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveUser) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestRemoveUser: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestRemoveUser) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseRemoveUser{}

	var dbReq db_local.User
	dbReq.Id = int64(request.Id)
	err := db_local.RemoveUserById(db_local.DB_LOCAL, &dbReq)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
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
