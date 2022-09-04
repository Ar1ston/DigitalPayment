package Requests

import (
	"DigitalPayment/Services/Publishers/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
	"gorm.io/gorm"
)

func init() {
	method := "GetPublisher"
	register_requests.Register(method, (*RequestGetPublisher)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestGetPublisher struct {
	Id uint64 `json:"id"`
}
type ResponseGetPublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Errno       uint64 `json:"errno"`
	Error       string `json:"error,omitempty"`
}

func (request *RequestGetPublisher) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestGetPublisher) Validation() []byte {
	isError := false
	rpl := ResponseGetPublisher{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION GetPublisher: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestGetPublisher) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseGetPublisher{}

	publisher, err := db_local.FindPublisherById(db_local.DB_LOCAL, map[string]interface{}{
		"id": &request.Id,
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
		rpl.Id = uint64(publisher.Id)
		rpl.Name = publisher.Name
		rpl.Description = publisher.Description
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
