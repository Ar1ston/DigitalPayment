package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"DigitalPayment/lib/register_requests"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	method := "GetAuthors"
	register_requests.Register(method, (*RequestGetAuthors)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetAuthors struct{}
type Author struct {
	Id          uint64 `json:"id"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseGetAuthors struct {
	Authors []Author
	Errno   uint64 `json:"errno"`
	Error   string `json:"error,omitempty"`
}

func (request *RequestGetAuthors) Decode(decReq []byte) *error {

	return nil
}
func (request *RequestGetAuthors) Validation() *error {

	return nil
}
func (request *RequestGetAuthors) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetAuthors{}

	authors, err := db_local.FindAuthors(db_local.DB_LOCAL, map[string]interface{}{})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range authors.Authors {
			rpl.Authors = append(rpl.Authors, Author{
				Id:          uint64(v.Id),
				FirstName:   v.FirstName,
				LastName:    v.LastName,
				Description: v.Description,
			})
		}
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
