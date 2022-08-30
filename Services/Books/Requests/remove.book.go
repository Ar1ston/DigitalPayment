package Requests

import (
	"DigitalPayment/Services/Books/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "RemoveBook"
	register_requests.Register(method, (*RequestRemoveBook)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestRemoveBook struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveBook struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveBook) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestRemoveBook) Validation() []byte {
	isError := false
	rpl := ResponseRemoveBook{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION RemoveBook: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestRemoveBook) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseRemoveBook{}

	var dbReq db_local.Book
	dbReq.Id = int64(request.Id)
	err := db_local.RemoveBookById(db_local.DB_LOCAL, &dbReq)

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
