package Requests

import (
	"DigitalPayment/Publishers/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func Hello() {}
func init() {
	method := "ChangePublisher"
	register_requests.Register(method, (*RequestChangePublisher)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestChangePublisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangePublisher) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestChangePublisher) Validation() []byte {
	isError := false
	rpl := ResponseChangePublisher{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION ChangePublisher: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestChangePublisher) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseChangePublisher{}

	updates := map[string]interface{}{}

	if request.Name != "" {
		updates["Name"] = request.Name
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	publisher, err := db_local.ChangePublisherById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(publisher.Id)
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
