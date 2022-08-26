package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
	"gorm.io/gorm"
)

func init() {
	method := "CheckUser"
	register_requests.Register(method, (*RequestCheckUser)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestCheckUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type ResponseCheckUser struct {
	Id    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Level uint64 `json:"level"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCheckUser) Decode(decReq []byte) *error {
	var rplBytes = bytes.NewBuffer(decReq)
	dec := gob.NewDecoder(rplBytes)
	err := dec.Decode(request)
	if err != nil {
		return &err
	}
	return nil
}

func (request *RequestCheckUser) Validation() *error {
	var err error
	if request.Login == "" {
		err = fmt.Errorf("%s", "Неверное поле Login в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCheckUser: %s\n", err.Error())
		return &err
	}
	if request.Password == "" {
		err = fmt.Errorf("%s", "Неверное поле password в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestCheckUser: %s\n", err.Error())
		return &err
	}
	return nil
}

func (request *RequestCheckUser) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseCheckUser{}

	user, err := db_local.FindUser(db_local.DB_LOCAL, map[string]interface{}{
		"login": request.Login,
	})
	if err == gorm.ErrRecordNotFound {
		rpl.Error = err.Error()
		rpl.Errno = 404
	}
	if err != gorm.ErrRecordNotFound && err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	}
	if err == nil {
		if user.Password != request.Password {
			rpl.Error = "Пароль неверный"
			rpl.Errno = 401
		} else {
			rpl.Id = uint64(user.Id)
			rpl.Name = user.Name
			rpl.Login = user.Login
			rpl.Level = uint64(user.Level)
			rpl.Errno = 0
		}

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
