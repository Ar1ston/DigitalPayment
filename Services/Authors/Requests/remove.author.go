package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "RemoveAuthor"
	register_requests.Register(method, (*RequestRemoveAuthor)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestRemoveAuthor struct {
	Id uint64 `json:"id"`
}
type ResponseRemoveAuthor struct {
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestRemoveAuthor) Decode(decReq []byte) *error {

	var rplBytes = bytes.NewBuffer(decReq)
	dec := gob.NewDecoder(rplBytes)
	err := dec.Decode(request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestRemoveAuthor) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestRemoveAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseRemoveAuthor{}

	var dbReq db_local.Author
	dbReq.Id = int64(request.Id)
	err := db_local.RemoveAuthorById(db_local.DB_LOCAL, &dbReq)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Errno = 0
	}
	fmt.Printf("RESPONSE: %+v\n", rpl)

	var rplBytes bytes.Buffer
	enc := gob.NewEncoder(&rplBytes)

	err = enc.Encode(rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes.Bytes(), nil
}
