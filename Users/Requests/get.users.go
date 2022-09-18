package Requests

import (
	"DigitalPayment/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "GetUsers"
	register_requests.Register(method, (*RequestGetUsers)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestGetUsers struct{}

type user struct {
	Id    uint64 `json:"id"`
	Login string `json:"name"`
	Level uint64 `json:"level"`
}
type ResponseGetUsers struct {
	Users []user
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestGetUsers) Decode(decReq []byte) *error {
	return nil
}
func (request *RequestGetUsers) Validation() []byte {
	return nil
}
func (request *RequestGetUsers) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseGetUsers{}

	users, err := db_local.FindUsers(db_local.DB_LOCAL, map[string]interface{}{})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range users.Users {
			rpl.Users = append(rpl.Users, user{
				Id:    uint64(v.Id),
				Login: v.Login,
				Level: uint64(v.Level),
			})
		}
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
