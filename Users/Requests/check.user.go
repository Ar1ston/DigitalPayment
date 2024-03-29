package Requests

import (
	"DigitalPayment/Users/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
	"gorm.io/gorm"
)

func init() {
	method := "CheckUser"
	register_requests.Register(method, (*RequestCheckUser)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
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
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}

func (request *RequestCheckUser) Validation() []byte {
	isError := false
	rpl := ResponseCheckUser{}
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
		logs.Logger.Errorf("ERROR VALIDATION CheckUser: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}

func (request *RequestCheckUser) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseCheckUser{}

	user, err := db_local.FindUser(db_local.DB_LOCAL, map[string]interface{}{
		"Login": request.Login,
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

		decryptPasswordIn, _ := crypt.Aes_decrypt(user.Password)
		decryptPasswordOut, _ := crypt.Aes_decrypt(request.Password)

		if decryptPasswordIn != decryptPasswordOut {
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
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil

}
