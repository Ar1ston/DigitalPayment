package Requests

import (
	"DigitalPayment/Books/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/logs"
	"DigitalPayment/lib/register_requests"
)

func init() {
	method := "GetBooks"
	register_requests.Register(method, (*RequestGetBooks)(nil))
	logs.Logger.Infof("Метод %s инициализирован!", method)
}

type RequestGetBooks struct{}
type Book struct {
	Id     int64
	Name   string
	Genre  string
	Author int64
}
type ResponseGetBooks struct {
	Books []Book
	Errno uint64 `json:"errno"`
	Error string `json:"error,omitempty"`
}

func (request *RequestGetBooks) Decode(decReq []byte) *error {
	return nil
}
func (request *RequestGetBooks) Validation() []byte {

	return nil
}
func (request *RequestGetBooks) Execute() ([]byte, *error) {
	logs.Logger.Infof("REQUEST: %+v", request)

	rpl := ResponseGetBooks{}

	books, err := db_local.FindBooks(db_local.DB_LOCAL, nil)

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range books.Books {
			rpl.Books = append(rpl.Books, Book{
				Id:     v.Id,
				Name:   v.Name,
				Genre:  v.Genre,
				Author: v.Author,
			})
		}
		rpl.Errno = 0
	}
	logs.Logger.Infof("RESPONSE: %+v", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
