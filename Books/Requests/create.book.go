package Requests

import (
	db_local2 "DigitalPayment/Books/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
	"time"
)

func init() {
	method := "CreateBook"
	register_requests.Register(method, (*RequestCreateBook)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestCreateBook struct {
	Name        string `json:"name,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Author      uint64 `json:"author,omitempty"`
	Publisher   uint64 `json:"publisher,omitempty"`
	Added_User  uint64 `json:"addedUser,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseCreateBook struct {
	Id    uint64 `json:"id"`
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestCreateBook) Decode(decReq []byte) *error {
	err := crypt.Gob_decrypt(decReq, request)
	if err != nil {
		return &err
	}
	return nil
}
func (request *RequestCreateBook) Validation() []byte {
	isError := false
	rpl := ResponseCreateBook{}
	if request.Name == "" {
		isError = true
		rpl.Errno = 409
		rpl.Error = "Error validation Name field in request"
	}
	if isError == false {
		return nil
	} else {
		logs.Logger.Errorf("ERROR VALIDATION CreateBook: %s", rpl.Error)
		encrypt, _ := crypt.Gob_encrypt(&rpl)
		return encrypt
	}
}
func (request *RequestCreateBook) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseCreateBook{}

	req := db_local2.Book{
		Name:        request.Name,
		Genre:       request.Genre,
		Author:      int64(request.Author),
		Publisher:   int64(request.Publisher),
		AddedUser:   int64(request.Added_User),
		AddedTime:   time.Now(),
		Description: request.Description,
	}

	book, err := db_local2.CreateBook(db_local2.DB_LOCAL, &req)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(book.Id)
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
