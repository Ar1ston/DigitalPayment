package Requests

import (
	"DigitalPayment/Services/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "RemoveUser"
	register_requests.Register(method, (*RequestRemoveUser)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestRemoveUser struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveUser) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestRemoveUser) Validation() []byte {
	isError := false
	rpl := ResponseRemoveUser{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Id field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION RemoveUser: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestRemoveUser) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

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
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
