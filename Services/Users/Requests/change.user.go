package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func Hello() {}
func init() {
	method := "ChangeUser"
	register_requests.Register(method, (*RequestChangeUser)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestChangeUser struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    uint64 `json:"level"`
}
type ResponseChangeUser struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeUser) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле Id в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestChangeUser: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestChangeUser) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseChangeUser{}

	updates := map[string]interface{}{}

	if request.Login != "" {
		updates["Login"] = request.Login

	}
	if request.Password != "" {
		updates["Password"] = request.Password

	}
	if request.Name != "" {
		updates["Name"] = request.Name
	}

	user, err := db_local.ChangeUserById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

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
