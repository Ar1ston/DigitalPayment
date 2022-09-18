package Requests

import (
	"DigitalPayment/Books/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func Hello() {}
func init() {
	method := "ChangeBook"
	register_requests.Register(method, (*RequestChangeBook)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestChangeBook struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseChangeBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestChangeBook) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestChangeBook) Validation() []byte {
	isError := false
	rpl := ResponseChangeBook{}
	if request.Id == 0 {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation ID field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION ChangeBook: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt

	}
}
func (request *RequestChangeBook) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseChangeBook{}

	updates := map[string]interface{}{}

	if request.Name != "" {
		updates["Name"] = request.Name
	}
	if request.Genre != "" {
		updates["Genre"] = request.Genre
	}
	if request.Author != 0 {
		updates["Author"] = request.Author
	}
	if request.Publisher != 0 {
		updates["Publisher"] = request.Publisher
	}
	if request.Description != "" {
		updates["Description"] = request.Description
	}

	books, err := db_local.ChangeBookById(db_local.DB_LOCAL, map[string]interface{}{"id": request.Id}, updates)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(books.Id)
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
