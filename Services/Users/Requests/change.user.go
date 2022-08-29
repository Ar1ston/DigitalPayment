package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
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

func (request *RequestChangeUser) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestChangeUser) Validation() []byte {
	isError := false
	rpl := ResponseChangeUser{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
		fmt.Printf("ERROR VALIDATION: %s\n", rpl.Error)
	}
	if isError == false {
		return nil
	} else {
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
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

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
