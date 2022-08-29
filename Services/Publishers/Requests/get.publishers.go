package Requests

import (
	"DigitalPayment/Services/Publishers/lib/db_local"
	"DigitalPayment/lib/crypt"
	"DigitalPayment/lib/register_requests"
	"fmt"
)

func init() {
	method := "GetPublishers"
	register_requests.Register(method, (*RequestGetPublishers)(nil))
	fmt.Printf("Метод %s инициализирован!\n", method)
}

type RequestGetPublishers struct{}
type Publisher struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ResponseGetPublishers struct {
	Publishers []Publisher `json:"publishers"`
	Errno      uint64      `json:"errno"`
	Error      string      `json:"error,omitempty"`
}

func (request *RequestGetPublishers) Decode(decReq []byte) *error {
	return nil
}
func (request *RequestGetPublishers) Validation() []byte {
	return nil
}
func (request *RequestGetPublishers) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetPublishers{}

	publishers, err := db_local.FindPublishers(db_local.DB_LOCAL, map[string]interface{}{})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		for _, v := range publishers.Publishers {
			rpl.Publishers = append(rpl.Publishers, Publisher{
				Id:          uint64(v.Id),
				Name:        v.Name,
				Description: v.Description,
			})
		}
		rpl.Errno = 0
	}
	fmt.Printf("RESPONSE: %+v\n", rpl)

	rplBytes, err := crypt.Gob_encrypt(&rpl)
	if err != nil {
		return nil, &err
	}

	return rplBytes, nil
}
