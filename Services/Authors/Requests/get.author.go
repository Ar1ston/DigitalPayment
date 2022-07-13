package Requests

import (
	"DigitalPayment/Services/Authors/lib/db_local"
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {

}

type RequestGetAuthor struct {
	Id uint64 `json:"id"`
}
type ResponseGetAuthor struct {
	Id          uint64 `json:"id" gorm:"id"`
	FirstName   string `json:"firstName,omitempty" gorm:"FirstName"`
	LastName    string `json:"lastName,omitempty" gorm:"LastName"`
	Description string `json:"description,omitempty" gorm:"Description"`
	Errno       uint64 `json:"errno" gorm:"-:all"`
	Error       string `json:"error,omitempty" gorm:"-:all"`
}

func (request *RequestGetAuthor) Validation() *error {
	var err error
	if request.Id == 0 {
		err = fmt.Errorf("%s", "Неверное поле ID в запросе")
		fmt.Printf("ОШИБКА ВАЛИДАЦИИ RequestGetAuthor: %s\n", err.Error())
		return &err
	}
	return nil
}
func (request *RequestGetAuthor) Execute() ([]byte, *error) {
	fmt.Printf("REQUEST: %+v\n", request)

	rpl := ResponseGetAuthor{}

	author, err := db_local.FindAuthorById(db_local.DB_LOCAL, map[string]interface{}{
		"id": request.Id,
	})

	if err != nil {
		rpl.Error = err.Error()
		rpl.Errno = 500
	} else {
		rpl.Id = uint64(author.Id)
		rpl.FirstName = author.FirstName
		rpl.LastName = author.LastName
		rpl.Description = author.Description
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
