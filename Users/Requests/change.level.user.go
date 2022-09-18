package Requests

import (
	"DigitalPayment/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "ChangeLevelUser"
	register_requests.Register(method, (*RequestChangeLevelUser)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestChangeLevelUser struct {
	Id    int64 `json:"id"`
	Level int8  `json:"level"`
}
type ResponseChangeLevelUser struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeLevelUser) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestChangeLevelUser) Validation() []byte {
	isError := false
	rpl := ResponseChangeLevelUser{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if request.Level > 1 || request.Level < -1 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Level field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION ChangeLevelUser: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestChangeLevelUser) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseChangeUser{}

	updates := map[string]interface{}{}

	user, err := db_local.FindUser(db_local.DB_LOCAL, map[string]interface{}{
		"id": request.Id,
	})

	if int8(user.Level)+request.Level > 3 || int8(user.Level)+request.Level < 1 {
		rpl.Error = "Error changing user level"
		rpl.Errno = 500
	} else {
		updates["Level"] = request.Level + int8(user.Level)

		user, err := db_local.ChangeUserById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

		if err != nil {
			rpl.Error = err.Error()
			rpl.Errno = 500
		} else {
			rpl.Id = uint64(user.Id)
			rpl.Errno = 0
		}
	}

	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
