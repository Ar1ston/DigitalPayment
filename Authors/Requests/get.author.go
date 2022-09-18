package Requests

import (
	"DigitalPayment/Authors/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
	"gorm.io/gorm"
)

func init() {
	method := "GetAuthor"
	register_requests.Register(method, (*RequestGetAuthor)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestGetAuthor struct {
	Id uint64 `json:"id"`
}
type ResponseGetAuthor struct {
	Id          uint64 `json:"id"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}

func (request *RequestGetAuthor) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestGetAuthor) Validation() []byte {
	isError := false
	rpl := ResponseGetAuthor{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION GetAuthor: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestGetAuthor) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseGetAuthor{}

	author, err := db_local.FindAuthorById(db_local.DB_LOCAL, map[string]interface{}{
		"id": request.Id,
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rpl.Error = err.Error()
			rpl.Errno = 404
		} else {
			rpl.Error = err.Error()
			rpl.Errno = 500
		}
	} else {
		rpl.Id = uint64(author.Id)
		rpl.FirstName = author.FirstName
		rpl.LastName = author.LastName
		rpl.Description = author.Description
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
