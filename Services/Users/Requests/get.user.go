package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "GetUser"
	register_requests.Register(method, (*RequestGetUser)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestGetUser struct {
	Id uint64 `json:"id"`
}
type ResponseGetUser struct {
	Id    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Level uint64 `json:"level"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestGetUser) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestGetUser) Validation() []byte {
	isError := false
	rpl := ResponseGetUser{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Id field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION GetUser: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestGetUser) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseGetUser{}

	user, err := db_local.FindUser(db_local.DB_LOCAL, map[string]interface{}{
		"id": request.Id,
	})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(user.Id)
		rpl.Name = user.Name
		rpl.Login = user.Login
		rpl.Level = uint64(user.Level)
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
