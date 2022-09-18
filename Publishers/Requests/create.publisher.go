package Requests

import (
	db_local2 "DigitalPayment/Publishers/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "CreatePublisher"
	register_requests.Register(method, (*RequestCreatePublisher)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestCreatePublisher struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreatePublisher struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreatePublisher) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreatePublisher) Validation() []byte {
	isError := false
	rpl := ResponseCreatePublisher{}
	if request.Name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Name field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION CreatePublisher: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestCreatePublisher) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseCreatePublisher{}

	req := db_local2.Publisher{
		Name:        request.Name,
		Description: request.Description,
	}

	publisher, err := db_local2.CreatePublisher(db_local2.DB_LOCAL, &req)

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
