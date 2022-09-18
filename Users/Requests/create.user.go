package Requests

import (
	db_local2 "DigitalPayment/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "CreateUser"
	register_requests.Register(method, (*RequestCreateUser)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
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
	}
	if request.Password == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Password field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION CreateUser: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestCreateUser) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseCreateUser{}

	req := db_local2.User{
		Login:    request.Login,
		Password: request.Password,
		Name:     request.Name,
	}

	user, err := db_local2.CreateUser(db_local2.DB_LOCAL, &req)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(user.Id)
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
